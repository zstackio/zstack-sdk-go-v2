// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateFiSecSecretResourcePool creates FiSecSecretResourcePool
func (cli *ZSClient) CreateFiSecSecretResourcePool(ctx context.Context, params param.CreateFiSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/secret-resource-pool/fiSec", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateFiSecSecretResourcePool updates FiSecSecretResourcePool
func (cli *ZSClient) UpdateFiSecSecretResourcePool(ctx context.Context, uuid string, params param.UpdateFiSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/secret-resource-pools/fiSec", uuid, "", map[string]interface{}{
		"updateFiSecSecretResourcePool": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
