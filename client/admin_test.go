package client

import (
	"fmt"
	"github.com/1uvu/fabric-sdk-client/types"
	"path/filepath"
	"testing"
)

func TestGetAdminClient(t *testing.T) {
	fmt.Println("testing admin client")

	var (
		basePath string = filepath.Join(
			"..",
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
		{"DISCOVERY_AS_LOCALHOST", "true"},
		{"TEST_IN_SHELL", "false"},
	}

	admin, err := GetAdminClient(params, envPairs...)

	if err != nil {
		t.Errorf("Failed to get admin client: %s\n", err)
	}

	err = testGetAppClient(admin)

	if err != nil {
		t.Error(err)
	}
}

func testGetAppClient(admin *AdminClient) error {
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
