package client

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/1uvu/fabric-sdk-client/types"
)

var app *AppClient

func TestGetAppClient(t *testing.T) {
	fmt.Println("testing app client")
	var (
		basePath string = filepath.Join(
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

	// 新建 app user 前要删除当前目录下的 wallet
	removeWallet()
	app2, err := GetAppClient("channel2", params, envPairs...)
	if err != nil {
		t.Errorf("Failed to get app client: %s", err)
	}

	// 全局赋值
	app = app2

	_, err = GetAppClient("channel12", params, envPairs...)
	if err != nil {
		t.Errorf("Failed to get app client: %s", err)
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
	_ = os.RemoveAll("./keystore/")
	_ = os.RemoveAll("./wallet/")
}
