package plugin

import (
	"context"
)

type Chainer interface {
	// AppPath returns the configured App's path.
	AppPath() string
	// ID returns the configured App's chain id.
	ID() (string, error)
	// ConfigPath returns the path to the App's config file.
	ConfigPath() string
	// RPCPublicAddress returns the configured App's rpc endpoint.
	RPCPublicAddress() (string, error)
}

// NewClientAPI creates a new app ClientAPI.
func NewClientAPI(c Chainer) ClientAPI {
	return clientAPI{chain: c}
}

type clientAPI struct {
	chain Chainer
}

func (api clientAPI) GetChainInfo(_ context.Context) (*ChainInfo, error) {
	chainID, err := api.chain.ID()
	if err != nil {
		return nil, err
	}
	appPath := api.chain.AppPath()
	configPath := api.chain.ConfigPath()
	rpc, err := api.chain.RPCPublicAddress()
	if err != nil {
		return nil, err
	}
	return &ChainInfo{
		ChainId:    chainID,
		AppPath:    appPath,
		ConfigPath: configPath,
		RpcAddress: rpc,
	}, nil
}