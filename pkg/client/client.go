package client

import (
	"fmt"
	"strings"
	"sync"

	"github.com/1uvu/fabric-sdk-client/pkg/config"
	"github.com/1uvu/fabric-sdk-client/pkg/sdk"
)

//
// 根据配置获取 Client 单例, 支持并发访问
//

var (
	apps   sync.Map
	admins sync.Map
)

// todo 修改为 abstract factory
func GetApp(channelID, clientConfigPath string) (*sdk.AppClient, error) {
	key := strings.Join([]string{channelID, clientConfigPath}, "-")
	if _, ok := apps.Load(key); !ok {
		app, err := newApp(channelID, clientConfigPath)
		if err != nil {
			return nil, err
		}
		apps.Store(key, app)
	}

	app, ok := apps.Load(key)
	if !ok {
		return nil, fmt.Errorf("failed to get app client of %s", channelID)
	}

	return app.(*sdk.AppClient), nil
}

func GetAdmin(clientConfigPath string) (*sdk.AdminClient, error) {
	key := clientConfigPath
	if _, ok := admins.Load(key); !ok {
		admin, err := newAdmin(clientConfigPath)
		if err != nil {
			return nil, err
		}

		admins.Store(key, admin)
	}

	admin, ok := admins.Load(key)
	if !ok {
		return nil, fmt.Errorf("failed to get admin client")
	}

	return admin.(*sdk.AdminClient), nil
}

func newApp(channelID, clientConfigPath string) (*sdk.AppClient, error) {

	conf, err := config.NewClientConfig(clientConfigPath)

	if err != nil {
		return nil, fmt.Errorf("failed to get app client: %s", err)
	}

	params := &conf.App.Params

	envPairs := conf.App.Envs

	app, err := sdk.GetAppClient(channelID, params, envPairs...)
	if err != nil {
		return nil, fmt.Errorf("failed to get app client: %s", err)
	}

	return app, nil
}

func newAdmin(clientConfigPath string) (*sdk.AdminClient, error) {

	conf, err := newClientConfig(clientConfigPath)

	if err != nil {
		return nil, fmt.Errorf("failed to get app client: %s", err)
	}

	params := &conf.Admin.Params

	envPairs := conf.Admin.Envs

	admin, err := sdk.GetAdminClient(params, envPairs...)
	if err != nil {
		return nil, fmt.Errorf("failed to get app client: %s", err)
	}

	return admin, nil
}
