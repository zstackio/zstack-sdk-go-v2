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
	err := cli.PutWithSpec("v1/hybrid/aliyun/vpc", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
