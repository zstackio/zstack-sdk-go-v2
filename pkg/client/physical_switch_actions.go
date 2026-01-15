// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPhysicalSwitch queries PhysicalSwitch list
func (cli *ZSClient) QueryPhysicalSwitch(params *param.QueryParam) ([]view.PhysicalSwitchInventoryView, error) {
	var resp []view.PhysicalSwitchInventoryView
	return resp, cli.List("v1/topo/physical-switches", params, &resp)
}

func (cli *ZSClient) GetPhysicalSwitch(uuid string) (*view.PhysicalSwitchInventoryView, error) {
	var resp view.PhysicalSwitchInventoryView
	if err := cli.Get("v1/topo/physical-switches", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePhysicalSwitch Pagination
func (cli *ZSClient) PagePhysicalSwitch(params *param.QueryParam) ([]view.PhysicalSwitchInventoryView, int, error) {
	var physicalSwitchs []view.PhysicalSwitchInventoryView
	total, err := cli.Page("v1/topo/physical-switches", params, &physicalSwitchs)
	return physicalSwitchs, total, err
}
