// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePluginSecretResourcePool creates PluginSecretResourcePool
func (cli *ZSClient) CreatePluginSecretResourcePool(ctx context.Context, params param.CreatePluginSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.Post(ctx, "v1/secret-resource-pool/plugin", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePluginSecretResourcePool updates PluginSecretResourcePool
func (cli *ZSClient) UpdatePluginSecretResourcePool(ctx context.Context, uuid string, params param.UpdatePluginSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/secret-resource-pool/plugin", uuid, "", map[string]interface{}{
		"updatePluginSecretResourcePool": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
