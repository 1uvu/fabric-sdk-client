package types

// 用于连接 Fabric 网络的参数
type AppParams struct {
	OrgName    string `json:"orgName" yaml:"orgName"`
	OrgMSP     string `json:"orgMSP" yaml:"orgMSP"`
	OrgHost    string `json:"orgHost" yaml:"orgHost"`
	OrgAdmin   string `json:"orgAdmin" yaml:"orgAdmin"`
	ConfigPath string `json:"configPath" yaml:"configPath"`
	OrgUser    string `json:"orgUser" yaml:"orgUser"`
	CredPath   string `json:"credPath" yaml:"credPath"`
	CertPath   string `json:"certPath" yaml:"certPath"`
}

// 用于连接 Fabric 网络的参数
type AdminParams struct {
	OrgName    string `json:"orgName" yaml:"orgName"`
	OrgMSP     string `json:"orgMSP" yaml:"orgMSP"`
	OrgHost    string `json:"orgHost" yaml:"orgHost"`
	OrgAdmin   string `json:"orgAdmin" yaml:"orgAdmin"`
	ConfigPath string `json:"configPath" yaml:"configPath"`
}
