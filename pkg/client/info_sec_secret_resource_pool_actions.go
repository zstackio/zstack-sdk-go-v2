// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateInfoSecSecretResourcePool creates InfoSecSecretResourcePool
func (cli *ZSClient) CreateInfoSecSecretResourcePool(params param.CreateInfoSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.CreateSecretResourcePoolEventView
	if err := cli.Post("v1/secret-resource-pool/infoSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateInfoSecSecretResourcePool updates InfoSecSecretResourcePool
func (cli *ZSClient) UpdateInfoSecSecretResourcePool(uuid string, params param.UpdateInfoSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.UpdateSecretResourcePoolEventView
	if err := cli.Put("v1/secret-resource-pools/infoSec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
