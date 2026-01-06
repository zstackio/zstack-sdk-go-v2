// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateHaiTaiSecretResourcePool creates HaiTaiSecretResourcePool
func (cli *ZSClient) CreateHaiTaiSecretResourcePool(params param.CreateHaiTaiSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.CreateSecretResourcePoolEventView
	if err := cli.Post("v1/secret-resource-pool/haitai", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
