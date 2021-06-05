package client

import (
	"fmt"
	"github.com/1uvu/fabric-sdk-client/types"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/events/deliverclient/seek"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
	"os"
)

type appClient struct {
	ChannelID string

	CC *channel.Client
	EC *event.Client
	LC *ledger.Client
}

type AdminClient struct {
	*types.AdminParams

	// sdk clients
	SDK *fabsdk.FabricSDK
	RC  *resmgmt.Client
	MC  *msp.Client

	ACs map[string]*appClient

	// for create channel
	// ChannelConfig string
	// OrdererID string

	valid bool
}

func InitAdminClient(params *types.AdminParams, pairs ...types.EnvPair) {
	adminParams = params
	for _, pair := range pairs {
		// such as "DISCOVERY_AS_LOCALHOST"
		_ = os.Setenv(pair.Key, pair.Val)
	}
}

var (
	adminParams *types.AdminParams
)

func GetAdminClient() (*AdminClient, error) {
	if adminParams.ConfigPath == "" {
		return nil, fmt.Errorf("Please init the params by call SerParams() function.")
	}

	sdk, err := fabsdk.New(config.FromFile(adminParams.ConfigPath))
	if err != nil {
		log.Panicf("failed to create fabric sdk: %s", err)
	}

	rcp := sdk.Context(fabsdk.WithUser(adminParams.OrgAdmin), fabsdk.WithOrg(adminParams.OrgName))
	rc, err := resmgmt.New(rcp)
	if err != nil {
		log.Panicf("failed to create resource client: %s", err)
	}

	mc, err := msp.New(sdk.Context())
	if err != nil {
		log.Panicf("failed to create msp client: %s", err)
	}

	admin := new(AdminClient)

	admin.AdminParams = adminParams
	admin.SDK = sdk
	admin.RC = rc
	admin.MC = mc
	admin.ACs = make(map[string]*appClient)
	admin.valid = true

	return admin, nil
}

func (admin *AdminClient) GetAppClient(channelID string) (*appClient, error) {

	if !admin.valid {
		return nil, fmt.Errorf("admin client is not valid now.\n")
	}

	if app, ok := admin.ACs[channelID]; ok {
		log.Printf("app client of %s has existed, return directly.\n", channelID)
		return app, nil
	}

	log.Printf("app client of %s do not existed, get it now.\n", channelID)

	app, err := admin.getAppClient(channelID)

	if err != nil {
		log.Fatalf("failed to app client %s.\n", channelID)
	}

	admin.ACs[channelID] = app

	return app, nil
}

func (admin *AdminClient) Exit() {
	admin.AdminParams = nil
	admin.SDK = nil
	admin.MC = nil
	admin.RC = nil

	for key := range admin.ACs {
		delete(admin.ACs, key)
	}
	admin.ACs = nil
	admin.valid = false

	adminParams = nil
}

func (admin *AdminClient) ExitApp(channelID string) error {
	if !admin.valid {
		return fmt.Errorf("admin client is not valid now.\n")
	}
	app, ok := admin.ACs[channelID]
	if !ok {
		return fmt.Errorf("app client of %s has existed, return directly.\n", channelID)
	}

	delete(admin.ACs, channelID)

	app.exit()

	return nil
}

func (admin *AdminClient) IsValid() bool {
	return admin.valid
}

func (admin *AdminClient) getAppClient(channelID string) (*appClient, error) {
	ccp := admin.SDK.ChannelContext(channelID, fabsdk.WithUser(admin.OrgAdmin))
	cc, err := channel.New(ccp)
	if err != nil {
		return nil, fmt.Errorf("failed to create channel client: %s", err)
	}

	ec, err := event.New(ccp, event.WithSeekType(seek.Newest))
	if err != nil {
		return nil, fmt.Errorf("failed to create event client: %s", err)
	}

	lc, err := ledger.New(ccp)
	if err != nil {
		return nil, fmt.Errorf("failed to create ledger client: %s", err)
	}

	app := new(appClient)
	app.ChannelID = channelID
	app.CC = cc
	app.EC = ec
	app.LC = lc

	return app, nil
}

func (app *appClient) exit() {
	app.ChannelID = ""
	app.CC = nil
	app.EC = nil
	app.LC = nil
}
