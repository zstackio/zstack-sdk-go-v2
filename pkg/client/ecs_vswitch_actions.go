// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEcsVSwitch updates EcsVSwitch
func (cli *ZSClient) UpdateEcsVSwitch(uuid string, params param.UpdateEcsVSwitchParam) (*view.EcsVSwitchInventoryView, error) {
	var resp view.UpdateEcsVSwitchEventView
	err := cli.PutWithSpec("v1/hybrid/aliyun/vswitch", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
