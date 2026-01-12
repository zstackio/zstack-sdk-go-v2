// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateMonitorGroup creates MonitorGroup
func (cli *ZSClient) CreateMonitorGroup(params param.CreateMonitorGroupParam) (*view.MonitorGroupInventoryView, error) {
	var resp view.CreateMonitorGroupEventView
	if err := cli.Post("v1/zwatch/monitorgroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteMonitorGroup deletes MonitorGroup
func (cli *ZSClient) DeleteMonitorGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitorgroups", uuid, string(deleteMode))
}
// QueryMonitorGroup queries MonitorGroup list
func (cli *ZSClient) QueryMonitorGroup(params *param.QueryParam) ([]view.MonitorGroupInventoryView, error) {
	var resp []view.MonitorGroupInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups", params, &resp)
}

func (cli *ZSClient) GetMonitorGroup(uuid string) (*view.MonitorGroupInventoryView, error) {
	var resp view.MonitorGroupInventoryView
	if err := cli.Get("v1/zwatch/monitorgroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateMonitorGroup updates MonitorGroup
func (cli *ZSClient) UpdateMonitorGroup(uuid string, params param.UpdateMonitorGroupParam) (*view.MonitorGroupInventoryView, error) {
	var resp view.UpdateMonitorGroupEventView
	if err := cli.Put("v1/zwatch/monitorgroups", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
