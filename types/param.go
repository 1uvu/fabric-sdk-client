package types

// 用于连接 Fabric 网络的参数
type AppParams struct {
	CredPath   string
	CertPath   string
	ConfigPath string
	OrgName    string
	OrgAdmin   string
	OrgUser    string
	OrgMSP     string
	OrgHost    string
}

// 用于连接 Fabric 网络的参数
type AdminParams struct {
	ConfigPath string
	OrgName    string
	OrgAdmin   string
	OrgMSP     string
	OrgHost    string
}

// chaincode invoke params
type InvokeParams struct {
	ChaincodeID string
	Fcn         string
	Args        [][]byte
	NeedSubmit  bool
	// for admin client
	Endpoints []string
}
