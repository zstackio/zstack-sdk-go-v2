// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePluginSecretResourcePool creates PluginSecretResourcePool
func (cli *ZSClient) CreatePluginSecretResourcePool(params param.CreatePluginSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.CreateSecretResourcePoolEventView
	if err := cli.Post("v1/secret-resource-pool/plugin", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdatePluginSecretResourcePool updates PluginSecretResourcePool
func (cli *ZSClient) UpdatePluginSecretResourcePool(uuid string, params param.UpdatePluginSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.UpdateSecretResourcePoolEventView
	if err := cli.Put("v1/secret-resource-pool/plugin/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
