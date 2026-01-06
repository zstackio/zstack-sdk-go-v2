// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateFiSecSecretResourcePool creates FiSecSecretResourcePool
func (cli *ZSClient) CreateFiSecSecretResourcePool(params param.CreateFiSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.CreateSecretResourcePoolEventView
	if err := cli.Post("v1/secret-resource-pool/fiSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateFiSecSecretResourcePool updates FiSecSecretResourcePool
func (cli *ZSClient) UpdateFiSecSecretResourcePool(uuid string, params param.UpdateFiSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.UpdateSecretResourcePoolEventView
	if err := cli.Put("v1/secret-resource-pools/fiSec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
