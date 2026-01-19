// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSecretResourcePool deletes SecretResourcePool
func (cli *ZSClient) DeleteSecretResourcePool(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/secret-resource-pool", uuid, string(deleteMode))
}
// UpdateSecretResourcePool updates SecretResourcePool
func (cli *ZSClient) UpdateSecretResourcePool(uuid string, params param.UpdateSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.Put("v1/secret-resource-pool", uuid, map[string]interface{}{
		"updateSecretResourcePool": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySecretResourcePool queries SecretResourcePool list
func (cli *ZSClient) QuerySecretResourcePool(params *param.QueryParam) ([]view.SecretResourcePoolInventoryView, error) {
	var resp []view.SecretResourcePoolInventoryView
	return resp, cli.List("v1/secret-resource-pools", params, &resp)
}

func (cli *ZSClient) GetSecretResourcePool(uuid string) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.SecretResourcePoolInventoryView
	if err := cli.Get("v1/secret-resource-pools", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSecretResourcePool Pagination
func (cli *ZSClient) PageSecretResourcePool(params *param.QueryParam) ([]view.SecretResourcePoolInventoryView, int, error) {
	var secretResourcePools []view.SecretResourcePoolInventoryView
	total, err := cli.Page("v1/secret-resource-pools", params, &secretResourcePools)
	return secretResourcePools, total, err
}
