// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEcsVpc updates EcsVpc
func (cli *ZSClient) UpdateEcsVpc(uuid string, params param.UpdateEcsVpcParam) (*view.EcsVpcInventoryView, error) {
	var resp view.UpdateEcsVpcEventView
	if err := cli.Put("v1/hybrid/aliyun/vpc/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
