// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSecretResourcePool deletes SecretResourcePool
func (cli *ZSClient) DeleteSecretResourcePool(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/secret-resource-pool", uuid, string(deleteMode))
}
// UpdateSecretResourcePool updates SecretResourcePool
func (cli *ZSClient) UpdateSecretResourcePool(ctx context.Context, uuid string, params param.UpdateSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/secret-resource-pool", uuid, "", map[string]interface{}{
		"updateSecretResourcePool": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySecretResourcePool queries SecretResourcePool list
func (cli *ZSClient) QuerySecretResourcePool(ctx context.Context, params *param.QueryParam) ([]view.SecretResourcePoolInventoryView, error) {
	var resp []view.SecretResourcePoolInventoryView
	return resp, cli.List(ctx, "v1/secret-resource-pools", params, &resp)
}

func (cli *ZSClient) GetSecretResourcePool(ctx context.Context, uuid string) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.SecretResourcePoolInventoryView
	if err := cli.Get(ctx, "v1/secret-resource-pools", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSecretResourcePool Pagination
func (cli *ZSClient) PageSecretResourcePool(ctx context.Context, params *param.QueryParam) ([]view.SecretResourcePoolInventoryView, int, error) {
	var secretResourcePools []view.SecretResourcePoolInventoryView
	total, err := cli.Page(ctx, "v1/secret-resource-pools", params, &secretResourcePools)
	return secretResourcePools, total, err
}
