// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSanSecSecretResourcePool creates SanSecSecretResourcePool
func (cli *ZSClient) CreateSanSecSecretResourcePool(params param.CreateSanSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.CreateSecretResourcePoolEventView
	if err := cli.Post("v1/secret-resource-pool/sanSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateSanSecSecretResourcePool updates SanSecSecretResourcePool
func (cli *ZSClient) UpdateSanSecSecretResourcePool(uuid string, params param.UpdateSanSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.UpdateSecretResourcePoolEventView
	if err := cli.Put("v1/secret-resource-pools/sanSec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
