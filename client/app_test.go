package client

import (
	"fmt"
	"github.com/1uvu/fabric-sdk-client/types"
	"os"
	"path/filepath"
	"testing"
)

func TestGetAppClient(t *testing.T) {
	fmt.Println("testing app client")

	var (
		basePath string = filepath.Join(
			"..",
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
		{"DISCOVERY_AS_LOCALHOST", "true"},
		{"TEST_IN_SHELL", "false"},
	}

	// 新建 app user 前要删除当前目录下的 wallet
	removeWallet()
	_, err := GetAppClient("channel2", params, envPairs...)
	if err != nil {
		t.Errorf("Failed to get app client: %s\n", err)
	}

	removeWallet()
	_, err = GetAppClient("channel12", params, envPairs...)
	if err != nil {
		t.Errorf("Failed to get app client: %s\n", err)
	}
}

func removeWallet() {
	_ = os.RemoveAll("./keystore/")
	_ = os.RemoveAll("./wallet/")
}
