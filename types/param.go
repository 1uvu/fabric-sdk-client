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
