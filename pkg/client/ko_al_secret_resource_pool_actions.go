// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateKoAlSecretResourcePool updates KoAlSecretResourcePool
func (cli *ZSClient) UpdateKoAlSecretResourcePool(uuid string, params param.UpdateKoAlSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.UpdateSecretResourcePoolEventView
	if err := cli.Put("v1/secret-resource-pools/koal/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateKoAlSecretResourcePool creates KoAlSecretResourcePool
func (cli *ZSClient) CreateKoAlSecretResourcePool(params param.CreateKoAlSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.CreateSecretResourcePoolEventView
	if err := cli.Post("v1/secret-resource-pool/koal", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
