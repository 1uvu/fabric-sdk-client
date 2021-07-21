package config

import (
	"log"
	"os"
	"testing"
)

func TestNewClientConfig(t *testing.T) {
	confPath := "./example_config/client-org1.yaml"

	conf, err := NewClientConfig(confPath)

	if err != nil {
		t.Error(err.Error())
	}

	log.Println(conf)
}

func TestExtendEnvPairs(t *testing.T) {
	envVal := os.Getenv("CRYPTO_CONFIG_BASE_PATH")

	if envVal == "" {
		t.Error("env pairs set error")
	}

	log.Println(envVal)
}
