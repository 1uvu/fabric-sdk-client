package types

// 用于连接 Fabric 网络的参数
type AppParams struct {
	OrgName    string `yaml:"orgName"`
	OrgMSP     string `yaml:"orgMSP"`
	OrgHost    string `yaml:"orgHost"`
	OrgAdmin   string `yaml:"orgAdmin"`
	ConfigPath string `yaml:"configPath"`
	OrgUser    string `yaml:"orgUser"`
	CredPath   string `yaml:"credPath"`
	CertPath   string `yaml:"certPath"`
}

// 用于连接 Fabric 网络的参数
type AdminParams struct {
	OrgName    string `yaml:"orgName"`
	OrgMSP     string `yaml:"orgMSP"`
	OrgHost    string `yaml:"orgHost"`
	OrgAdmin   string `yaml:"orgAdmin"`
	ConfigPath string `yaml:"configPath"`
}
