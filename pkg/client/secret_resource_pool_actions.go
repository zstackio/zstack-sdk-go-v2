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
	return cli.DeleteWithSpec("v1/secret-resource-pool", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// UpdateSecretResourcePool updates SecretResourcePool
func (cli *ZSClient) UpdateSecretResourcePool(uuid string, params param.UpdateSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.UpdateSecretResourcePoolEventView
	err := cli.PutWithSpec("v1/secret-resource-pool", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
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
