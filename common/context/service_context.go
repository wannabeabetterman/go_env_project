package context

import (
	"alger/config"
)

type ApiServiceContext struct {
	Config config.Config
	//UserRpc  user.UserService
	//AssetRpc asset.AssetService
	//Metadata metadata.MetadataService
	//DataAuth map[int64]dto.DataAuthDto
}

func NewApiServiceContext(c config.Config) *ApiServiceContext {
	return &ApiServiceContext{
		Config: c,
		//UserRpc:  user.NewUserService(zrpc.MustNewClient(c.Api.UserRpc)),
		//AssetRpc: asset.NewAssetService(zrpc.MustNewClient(c.Api.AssetRpc)),
		//Metadata: metadata.NewMetadataService(zrpc.MustNewClient(c.Api.MetadataRpc)),
		//DataAuth: make(map[int64]dto.DataAuthDto),
	}
}
