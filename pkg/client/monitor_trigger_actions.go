// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMonitorTrigger queries MonitorTrigger list
func (cli *ZSClient) QueryMonitorTrigger(params *param.QueryParam) ([]view.MonitorTriggerInventoryView, error) {
	var resp []view.MonitorTriggerInventoryView
	return resp, cli.List("v1/monitoring/triggers", params, &resp)
}

func (cli *ZSClient) GetMonitorTrigger(uuid string) (*view.MonitorTriggerInventoryView, error) {
	var resp view.MonitorTriggerInventoryView
	if err := cli.Get("v1/monitoring/triggers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateMonitorTrigger creates MonitorTrigger
func (cli *ZSClient) CreateMonitorTrigger(params param.CreateMonitorTriggerParam) (*view.MonitorTriggerInventoryView, error) {
	var resp view.CreateMonitorTriggerEventView
	if err := cli.Post("v1/monitoring/triggers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteMonitorTriggerAction deletes MonitorTrigger
func (cli *ZSClient) DeleteMonitorTriggerAction(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/trigger-actions", uuid, string(deleteMode))
}
// QueryMonitorTriggerAction queries MonitorTrigger list
func (cli *ZSClient) QueryMonitorTriggerAction(params *param.QueryParam) ([]view.MonitorTriggerActionInventoryView, error) {
	var resp []view.MonitorTriggerActionInventoryView
	return resp, cli.List("v1/monitoring/trigger-actions", params, &resp)
}

func (cli *ZSClient) GetMonitorTriggerAction(uuid string) (*view.MonitorTriggerActionInventoryView, error) {
	var resp view.MonitorTriggerActionInventoryView
	if err := cli.Get("v1/monitoring/trigger-actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteMonitorTrigger deletes MonitorTrigger
func (cli *ZSClient) DeleteMonitorTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/triggers", uuid, string(deleteMode))
}
// UpdateMonitorTrigger updates MonitorTrigger
func (cli *ZSClient) UpdateMonitorTrigger(uuid string, params param.UpdateMonitorTriggerParam) (*view.MonitorTriggerInventoryView, error) {
	var resp view.UpdateMonitorTriggerEventView
	if err := cli.Put("v1/monitoring/triggers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
