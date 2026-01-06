// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSecretResourcePool deletes SecretResourcePool
func (cli *ZSClient) DeleteSecretResourcePool(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/secret-resource-pool/{uuid}", uuid, string(deleteMode))
}
// UpdateSecretResourcePool updates SecretResourcePool
func (cli *ZSClient) UpdateSecretResourcePool(uuid string, params param.UpdateSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.UpdateSecretResourcePoolEventView
	if err := cli.Put("v1/secret-resource-pool/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySecretResourcePool queries SecretResourcePool list
func (cli *ZSClient) QuerySecretResourcePool(params *param.QueryParam) ([]view.SecretResourcePoolInventoryView, error) {
	var resp []view.SecretResourcePoolInventoryView
	return resp, cli.List("v1/secret-resource-pools", params, &resp)
}
