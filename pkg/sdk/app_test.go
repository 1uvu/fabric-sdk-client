package sdk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/1uvu/fabric-sdk-client/pkg/types"
)

var app *AppClient

func TestGetAppClient(t *testing.T) {
	fmt.Println("testing app client")

	// 首先要清空 wallet, 否则会出错
	removeWallet()

	var (
		basePath string = filepath.Join(
			"..",
			"..",
			"..",
			"Fabric-Demo",
			"network",
			"orgs",
		)
		orgName    string = "Org2"
		orgMSP     string = "Org2MSP"
		orgHost    string = "org2.example.com"
		configName string = "app-org2.yaml"
		orgUser    string = "User1"
		orgAdmin   string = "Admin"
	)

	credPath := filepath.Join(
		basePath,
		"peerOrganizations",
		orgHost,
		"users",
		fmt.Sprintf("%s@%s", orgUser, orgHost),
		"msp",
	)
	certPath := filepath.Join(
		credPath,
		"signcerts",
		fmt.Sprintf("%s@%s-cert.pem", orgUser, orgHost),
	)
	configPath := filepath.Join(
		basePath,
		"app",
		configName,
	)
	params := &types.AppParams{
		CredPath:   credPath,
		CertPath:   certPath,
		ConfigPath: configPath,
		OrgMSP:     orgMSP,
		OrgName:    orgName,
		OrgAdmin:   orgAdmin,
		OrgUser:    orgUser,
		OrgHost:    orgHost,
	}

	envPairs := []types.EnvPair{
		{Key: "DISCOVERY_AS_LOCALHOST", Val: "true"},
		{Key: "TEST_IN_SHELL", Val: "false"},
	}

	app2, err := GetAppClient("channel2", params, envPairs...)
	if err != nil {
		t.Errorf("Failed to get app2 client: %s", err)
	}

	// 全局赋值
	app = app2

	_, err = GetAppClient("channel12", params, envPairs...)
	if err != nil {
		t.Errorf("Failed to get app12 client: %s", err)
	}

	params.OrgName = "Org1"
	params.OrgMSP = "Org1MSP"
	params.OrgHost = "org1.example.com"
	configName = "app-org1.yaml"

	params.CredPath = filepath.Join(
		basePath,
		"peerOrganizations",
		params.OrgHost,
		"users",
		fmt.Sprintf("%s@%s", orgUser, params.OrgHost),
		"msp",
	)
	params.CertPath = filepath.Join(
		params.CredPath,
		"signcerts",
		fmt.Sprintf("%s@%s-cert.pem", orgUser, params.OrgHost),
	)
	params.ConfigPath = filepath.Join(
		basePath,
		"app",
		configName,
	)

	_, err = GetAppClient("channel1", params, envPairs...)
	if err != nil {
		t.Errorf("Failed to get app1 client: %s", err)
	}
}

func TestInvokeChaincode(t *testing.T) {
	// 已 app2 为例, 测试
	request := &types.InvokeRequest{
		ChaincodeID: "patient",
		Fcn:         "Query",
		Args:        []string{"h1"},
		NeedSubmit:  false,
	}

	resp, err := app.InvokeChaincode(request)

	if err != nil {
		t.Errorf("invoke test failed with error: %s", err)
	}

	log.Println("get the response as follows.")
	log.Println("payload: ", string(resp.Payload))
	log.Println("tx info: ", resp.TransactionInfo)
	log.Println("status code: ", resp.ChaincodeStatus)
}

func removeWallet() {
	_ = os.Remove("./keystore")
	_ = os.Remove("./wallet")
}
