package client

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/1uvu/fabric-sdk-client/types"
)

var (
	admin *AdminClient
	_app  *appClient
	txid  string
)

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

	log.Println("timesleep 3s to wait block broadcast.")
	time.Sleep(3 * time.Second)

	err = testGetBlockInfoByTxID(txid)

	if err != nil {
		t.Error(err)
	}

	err = testQueryChannelInfo()

	if err != nil {
		t.Error(err)
	}

}

func testGetAppClient() error {
	// 获取通道的 app client
	app2, err := admin.GetAppClient("channel2")

	_app = app2

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
		// "peer0.org1.example.com",
		"peer0.org2.example.com", // 根据背书策略来设置，也可置为空，使用默认设置
		// "peer0.org3.example.com",
	}
	request := &types.InvokeRequest{
		ChaincodeID: "patient",
		Fcn:         "Query",
		Args:        []string{"h1"},
		NeedSubmit:  true,
		Endpoints:   endpoints,
	}

	app, _ := admin.GetAppClient("channel2")
	resp, err := app.InvokeChaincode(request)

	if err != nil {
		return fmt.Errorf("invoke test failed with error: %s", err)
	}

	txid = resp.TransactionInfo.TransactionID

	log.Println("get the response as follows.")
	log.Println("payload: ", string(resp.Payload))
	log.Println("tx info: ", resp.TransactionInfo)
	log.Println("status code: ", resp.ChaincodeStatus)

	return nil
}

func testQueryChannelInfo() error {
	info, err := _app.QueryChannelInfo()

	if err != nil {
		return err
	}

	log.Println("channel info: ", info)

	return nil
}

func testGetBlockInfoByTxID(txid string) error {
	log.Println("txid: ", txid)
	info, err := _app.QueryBlockInfoByTxID(txid)

	if err != nil {
		return err
	}

	log.Println("block info: ", info)

	return nil
}
