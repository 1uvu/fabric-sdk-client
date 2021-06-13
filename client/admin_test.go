package client

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"github.com/1uvu/fabric-sdk-client/types"
)

var admin *AdminClient

func TestGetAdminClient(t *testing.T) {
	fmt.Println("testing admin client")

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
		configName string = "admin-org2.yaml"
		orgAdmin   string = "Admin"
	)

	configPath := filepath.Join(
		basePath,
		"admin",
		configName,
	)
	params := &types.AdminParams{
		ConfigPath: configPath,
		OrgMSP:     orgMSP,
		OrgName:    orgName,
		OrgAdmin:   orgAdmin,
		OrgHost:    orgHost,
	}

	envPairs := []types.EnvPair{
		{Key: "DISCOVERY_AS_LOCALHOST", Val: "true"},
		{Key: "TEST_IN_SHELL", Val: "false"},
	}

	admin2, err := GetAdminClient(params, envPairs...)

	if err != nil {
		t.Errorf("Failed to get admin client: %s", err)
	}

	// 全局赋值
	admin = admin2

	err = testGetAppClient()

	if err != nil {
		t.Error(err)
	}

	err = testInvokeChaincode()

	if err != nil {
		t.Error(err)
	}
}

func testGetAppClient() error {
	// 获取通道的 app client
	app2, err := admin.GetAppClient("channel2")

	if err != nil {
		return err
	}

	// 这里获取到的 _app2 与 app2 是同一个 app client
	_app2, err := admin.GetAppClient("channel2")

	if err != nil {
		return err
	}

	if app2 != _app2 {
		return fmt.Errorf("app2 not equal _app2")
	}

	return nil
}

func testInvokeChaincode() error {
	// 已 app2 为例, 测试
	endpoints := []string{
		"peer0.org1.example.com",
		"peer0.org2.example.com", // 三者至少包括两个即可
		"peer0.org3.example.com",
	}
	params := &types.InvokeParams{
		ChaincodeID: "patient",
		Fcn:         "Query",
		Args:        [][]byte{[]byte("h1")},
		NeedSubmit:  false,
		Endpoints:   endpoints,
	}

	app, _ := admin.GetAppClient("channel2")
	result, err := app.InvokeChaincode(params)

	if err != nil {
		return fmt.Errorf("invoke test failed with error: %s", err)
	}

	log.Println(string(result))
	return nil
}
