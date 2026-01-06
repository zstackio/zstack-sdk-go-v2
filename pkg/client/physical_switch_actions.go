// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPhysicalSwitch queries PhysicalSwitch list
func (cli *ZSClient) QueryPhysicalSwitch(params *param.QueryParam) ([]view.PhysicalSwitchInventoryView, error) {
	var resp []view.PhysicalSwitchInventoryView
	return resp, cli.List("v1/topo/physical-switches", params, &resp)
}
