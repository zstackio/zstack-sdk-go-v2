// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateKoAlSecretResourcePool updates KoAlSecretResourcePool
func (cli *ZSClient) UpdateKoAlSecretResourcePool(ctx context.Context, uuid string, params param.UpdateKoAlSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/secret-resource-pools/koal", uuid, "", map[string]interface{}{
		"updateKoAlSecretResourcePool": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateKoAlSecretResourcePool creates KoAlSecretResourcePool
func (cli *ZSClient) CreateKoAlSecretResourcePool(ctx context.Context, params param.CreateKoAlSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/secret-resource-pool/koal", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
