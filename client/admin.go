package client

import (
	"fmt"
	"log"
	"os"

	"github.com/1uvu/fabric-sdk-client/types"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/events/deliverclient/seek"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type appClient struct {
	ChannelID string

	metadata *types.AdminParams

	CC *channel.Client
	EC *event.Client
	LC *ledger.Client
}

type AdminClient struct {
	metadata *types.AdminParams

	// sdk clients
	SDK *fabsdk.FabricSDK
	RC  *resmgmt.Client
	MC  *msp.Client

	ACs map[string]*appClient

	// for create channel
	// ChannelConfig string
	// OrdererID string
}

func GetAdminClient(params *types.AdminParams, envPairs ...types.EnvPair) (*AdminClient, error) {
	// todo: params 检查合法性，如文件是否存在

	for _, pair := range envPairs {
		// such as "DISCOVERY_AS_LOCALHOST"
		_ = os.Setenv(pair.Key, pair.Val)
	}

	sdk, err := fabsdk.New(config.FromFile(params.ConfigPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create fabric sdk: %s", err)
	}

	rcp := sdk.Context(fabsdk.WithUser(params.OrgAdmin), fabsdk.WithOrg(params.OrgName))
	rc, err := resmgmt.New(rcp)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource client: %s", err)
	}

	mc, err := msp.New(sdk.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to create msp client: %s", err)
	}

	admin := new(AdminClient)

	admin.metadata = params
	admin.SDK = sdk
	admin.RC = rc
	admin.MC = mc
	admin.ACs = make(map[string]*appClient)

	return admin, nil
}

func (admin *AdminClient) GetAppClient(channelID string) (*appClient, error) {

	if app, ok := admin.ACs[channelID]; ok {
		log.Printf("app client of %s has existed, return directly", channelID)
		return app, nil
	}

	log.Printf("app client of %s do not existed, get it now", channelID)

	app, err := admin.getAppClient(channelID)

	if err != nil {
		return nil, fmt.Errorf("failed to get app client with error %s", err)
	}

	admin.ACs[channelID] = app

	return app, nil
}

func (admin *AdminClient) getAppClient(channelID string) (*appClient, error) {
	ccp := admin.SDK.ChannelContext(channelID, fabsdk.WithUser(admin.metadata.OrgAdmin))
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
	app.metadata = admin.metadata
	app.CC = cc
	app.EC = ec
	app.LC = lc

	return app, nil
}

// 这是一个非常简单的封装, 如需定义更多参数, 请直接使用 client 按照官方 sdk 定制
func (app *appClient) InvokeChaincode(request *types.InvokeRequest) (response *types.InvokeResponse, err error) {
	args := make([][]byte, len(request.Args))
	for i, a := range request.Args {
		args[i] = []byte(a)
	}

	req := channel.Request{
		ChaincodeID: request.ChaincodeID,
		Fcn:         request.Fcn,
		Args:        args,
	}

	reqPeers := channel.WithTargetEndpoints(request.Endpoints...)

	resp := *new(channel.Response)
	if request.NeedSubmit {
		resp, err = app.CC.Execute(req, reqPeers)
	} else {
		resp, err = app.CC.Query(req)
	}

	return &types.InvokeResponse{
		Payload:         resp.Payload,
		TransactionInfo: types.NewTransactionInfo(string(resp.TransactionID), app.metadata.OrgName),
		ChaincodeStatus: resp.ChaincodeStatus,
	}, err
}

func (app *appClient) QueryBlockInfoByTxID(txid string) (*types.BlockInfo, error) {
	block, err := app.LC.QueryBlockByTxID(fab.TransactionID(txid))

	if err != nil {
		return nil, err
	}

	return &types.BlockInfo{
		BlockNumber:  block.Header.Number,
		DataHash:     block.Header.DataHash,
		PreviousHash: block.Header.PreviousHash,
	}, nil
}

func (app *appClient) QueryChannelInfo() (*types.ChannelInfo, error) {
	channelCfg, err := app.LC.QueryConfig()

	if err != nil {
		return nil, err
	}

	chainInfo, err := app.LC.QueryInfo()

	if err != nil {
		return nil, err
	}

	return &types.ChannelInfo{
		ChannelID:         channelCfg.ID(),
		Height:            chainInfo.BCI.Height,
		CurrentBlockHash:  chainInfo.BCI.CurrentBlockHash,
		PreviousBlockHash: chainInfo.BCI.PreviousBlockHash,
		Orderers:          channelCfg.Orderers(),
		Endorser:          chainInfo.Endorser,
		Status:            chainInfo.Status,
	}, nil
}
