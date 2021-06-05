package client

import (
	"errors"
	"fmt"
	"github.com/1uvu/fabric-sdk-client/types"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type AppClient struct {
	*types.AppParams
	*gateway.Network
	valid bool
}

var (
	appParams *types.AppParams
)

func InitAppClient(params *types.AppParams, pairs ...types.EnvPair) {
	appParams = params
	for _, pair := range pairs {
		// such as "DISCOVERY_AS_LOCALHOST"
		_ = os.Setenv(pair.Key, pair.Val)
	}
}

func GetAppClient(channelID string) (*AppClient, error) {
	if appParams == nil {
		return nil, fmt.Errorf("please init the %s client firstly by call `InitAppClient()`.\n", channelID)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s\n", err)
			os.Exit(1)
		}
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(appParams.ConfigPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()

	network, err := gw.GetNetwork(channelID)
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}

	return &AppClient{AppParams: appParams, Network: network, valid: true}, nil
}

func (app *AppClient) IsValid() bool {
	return app.valid
}

func (app *AppClient) Exit() {
	app.Network = nil
	app.AppParams = nil
	app.valid = false

	appParams = nil
}

func populateWallet(wallet *gateway.Wallet) error {
	// todo run.sh 给予警告
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(appParams.CertPath))
	if err != nil {
		return err
	}

	// there's a single file in this dir containing the private key
	keyDir := filepath.Join(appParams.CredPath, "keystore")
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file\n")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(appParams.OrgMSP, string(cert), string(key))

	err = wallet.Put("appUser", identity)
	if err != nil {
		return err
	}
	return nil
}
