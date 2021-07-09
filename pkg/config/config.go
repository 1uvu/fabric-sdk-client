package config

import (
	"io/ioutil"
	"os"

	"github.com/1uvu/fabric-sdk-client/pkg/types"

	"gopkg.in/yaml.v2"
)

func NewClientConfig(confPath string) (*types.ClientConfig, error) {

	conf, err := getClientConfig(confPath)

	if err != nil {
		return nil, err
	}

	return conf, nil
}

func getClientConfig(confPath string) (*types.ClientConfig, error) {
	conf := new(types.ClientConfig)
	confFile, err := ioutil.ReadFile(confPath)

	if err != nil {
		return nil, err
	}

	extendEnvFile, err := extendEnvPairs(confFile)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(extendEnvFile, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func extendEnvPairs(confFile []byte) (extendEnvFile []byte, err error) {
	// 这里是为了使 global envs 生效
	_conf := new(types.ClientConfig)

	err = yaml.Unmarshal(confFile, _conf)

	if err != nil {
		return nil, err
	}

	for _, envPair := range _conf.Global {
		os.Setenv(envPair.Key, envPair.Val)
	}

	extendEnvFile = []byte(os.ExpandEnv(string(confFile)))

	return extendEnvFile, nil
}
