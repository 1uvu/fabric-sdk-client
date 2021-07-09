package sdk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/1uvu/fabric-sdk-client/pkg/types"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type AppClient struct {
	metadata *types.AppParams
	*gateway.Network
}

func GetAppClient(channelID string, params *types.AppParams, envPairs ...types.EnvPair) (*AppClient, error) {
	// todo: params 检查合法性，如文件是否存在

	for _, pair := range envPairs {
		// such as "DISCOVERY_AS_LOCALHOST"
		_ = os.Setenv(pair.Key, pair.Val)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser" + params.OrgName) {
		err = populateWallet(wallet, params)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s", err)
			os.Exit(1)
		}
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(params.ConfigPath))),
		gateway.WithIdentity(wallet, "appUser"+params.OrgName),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s", err)
		os.Exit(1)
	}
	defer gw.Close()

	network, err := gw.GetNetwork(channelID)
	if err != nil {
		fmt.Printf("Failed to get network: %s", err)
		os.Exit(1)
	}

	return &AppClient{metadata: params, Network: network}, nil
}

// 通过 AppClient 调用链码无法获取 txID 和 statusCode (官方库里面隐藏了)
// 因为 AppClient 本质上是一个轻量级的 CC Client
func (app *AppClient) InvokeChaincode(request *types.InvokeRequest) (response *types.InvokeResponse, err error) {
	contract := app.GetContract(request.ChaincodeID)

	var result []byte

	if request.NeedSubmit {
		result, err = contract.SubmitTransaction(request.Fcn, request.Args...)
	} else {
		result, err = contract.EvaluateTransaction(request.Fcn, request.Args...)
	}

	return &types.InvokeResponse{
		Payload:         result,
		TransactionInfo: types.NewTransactionInfo("", app.metadata.OrgName),
	}, err
}

func populateWallet(wallet *gateway.Wallet, params *types.AppParams) error {
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(params.CertPath))
	if err != nil {
		return err
	}

	// there's a single file in this dir containing the private key
	keyDir := filepath.Join(params.CredPath, "keystore")
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(params.OrgMSP, string(cert), string(key))

	err = wallet.Put("appUser"+params.OrgName, identity)
	if err != nil {
		return err
	}
	return nil
}
