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
	resp := view.EcsVSwitchInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/vswitch", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
