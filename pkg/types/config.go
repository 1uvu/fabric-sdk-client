package types

//
// 读入 *.yaml 配置文件
//

type ClientConfig struct {
	Global []GlobalEnvPair `json:"global" yaml:"global,flow"`
	Admin  AdminConfig     `json:"admin" yaml:"admin"`
	App    AppConfig       `json:"app" yaml:"app"`
}

type AdminConfig struct {
	Params AdminParams `json:"params" yaml:"params"`
	Envs   []EnvPair   `json:"envs" yaml:"envs,flow"`
}

type AppConfig struct {
	Params AppParams `json:"params" yaml:"params"`
	Envs   []EnvPair `json:"envs" yaml:"envs,flow"`
}

type GlobalEnvPair struct {
	Key string `json:"key" yaml:"key"`
	Val string `json:"val" yaml:"val"`
}
