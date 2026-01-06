// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateFlkSecSecretResourcePool updates FlkSecSecretResourcePool
func (cli *ZSClient) UpdateFlkSecSecretResourcePool(uuid string, params param.UpdateFlkSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.UpdateSecretResourcePoolEventView
	if err := cli.Put("v1/secret-resource-pools/flkSec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateFlkSecSecretResourcePool creates FlkSecSecretResourcePool
func (cli *ZSClient) CreateFlkSecSecretResourcePool(params param.CreateFlkSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.CreateSecretResourcePoolEventView
	if err := cli.Post("v1/secret-resource-pool/flkSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
