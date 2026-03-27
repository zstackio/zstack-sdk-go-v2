// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSanSecSecretResourcePool creates SanSecSecretResourcePool
func (cli *ZSClient) CreateSanSecSecretResourcePool(ctx context.Context, params param.CreateSanSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.Post(ctx, "v1/secret-resource-pool/sanSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSanSecSecretResourcePool updates SanSecSecretResourcePool
func (cli *ZSClient) UpdateSanSecSecretResourcePool(ctx context.Context, uuid string, params param.UpdateSanSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/secret-resource-pools/sanSec", uuid, "", map[string]interface{}{
		"updateSanSecSecretResourcePool": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
