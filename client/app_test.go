package client

//
//import (
//	"encoding/json"
//	"fmt"
//	"log"
//	"os"
//	"path/filepath"
//	"testing"
//
//	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
//)
//
//// 把 run sh 弄到代码里面
//
////
//// 定义全局变量
////
//
//var (
//	basePath string = filepath.Join(
//		"..",
//		"network",
//		"orgs",
//	)
//	orgName    string = "Org2"
//	orgMSP     string = "Org2MSP"
//	orgHost    string = "org2.example.com"
//	configName string = ""
//	orgUser    string = "User1"
//	orgAdmin   string = "Admin"
//)
//
//func Test(t *testing.T) {
//	appTest()
//	adminTest()
//}
//
//func appTest() {
//	fmt.Println("testing app client")
//	configName = "app-org2.yaml"
//
//	SetAppEnv("true")
//
//	credPath := filepath.Join(
//		basePath,
//		"peerOrganizations",
//		orgHost,
//		"users",
//		fmt.Sprintf("%s@%s", orgUser, orgHost),
//		"msp",
//	)
//	certPath := filepath.Join(
//		credPath,
//		"signcerts",
//		fmt.Sprintf("%s@%s-cert.pem", orgUser, orgHost),
//	)
//	configPath := filepath.Join(
//		basePath,
//		"app",
//		configName,
//	)
//	params := AppParams{ // todo 这里是指针，所以，初始化时直接变更即可
//		CredPath:   credPath,
//		CertPath:   certPath,
//		ConfigPath: configPath,
//		OrgMSP:     orgMSP,
//		OrgName:    orgName,
//		OrgAdmin:   orgAdmin,
//		OrgUser:    orgUser,
//		OrgHost:    orgHost,
//	}
//	SetAppParams(&params)
//
//	app2, err := GetAppClient("channel2")
//	if err != nil {
//		fmt.Printf("Failed to get app client: %s\n", err)
//		os.Exit(1)
//	}
//
//	app12, err := GetAppClient("channel12")
//	if err != nil {
//		fmt.Printf("Failed to get app client: %s\n", err)
//		os.Exit(1)
//	}
//
//	hid := "h3"
//	patient := structures.NewPatientInHIB(
//		[]string{
//			"ZJH-3",
//			"female",
//			"2020-10-10",
//			"abcdefghijklmnop",
//			"151",
//			"CQU",
//			"NMG",
//			"6674-1231-1000",
//		},
//	)
//
//	patientChaincode := app2.GetContract("patient")
//
//	result, err := patientChaincode.SubmitTransaction("register", []string{hid, patient.String()}...)
//	if err != nil {
//		fmt.Printf("Failed to submit transaction: %s\n", err)
//	}
//
//	result, err = patientChaincode.EvaluateTransaction("makeDigest", []string{hid}...)
//	if err != nil {
//		fmt.Printf("Failed to evaluate transaction: %s\n", err)
//	}
//	digest := new(DigestResult)
//	_ = json.Unmarshal(result, digest)
//	log.Println(digest.Digest)
//
//	bridgeChaincode := app12.GetContract("bridge")
//
//	_, err = bridgeChaincode.SubmitTransaction("register", []string{hid, digest.Digest}...)
//	if err != nil {
//		fmt.Printf("Failed to evaluate transaction: %s\n", err)
//	}
//}
//
