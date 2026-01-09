// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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
