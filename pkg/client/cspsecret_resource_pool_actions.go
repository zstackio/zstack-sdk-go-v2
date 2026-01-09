// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateCSPSecretResourcePool updates CSPSecretResourcePool
func (cli *ZSClient) UpdateCSPSecretResourcePool(uuid string, params param.UpdateCSPSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.UpdateSecretResourcePoolEventView
	if err := cli.Put("v1/secret-resource-pools/csp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateCSPSecretResourcePool creates CSPSecretResourcePool
func (cli *ZSClient) CreateCSPSecretResourcePool(params param.CreateCSPSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.CreateSecretResourcePoolEventView
	if err := cli.Post("v1/secret-resource-pool/csp", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
