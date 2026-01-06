// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySchedulerTrigger queries SchedulerTrigger list
func (cli *ZSClient) QuerySchedulerTrigger(params *param.QueryParam) ([]view.SchedulerTriggerInventoryView, error) {
	var resp []view.SchedulerTriggerInventoryView
	return resp, cli.List("v1/scheduler/triggers", params, &resp)
}
// UpdateSchedulerTrigger updates SchedulerTrigger
func (cli *ZSClient) UpdateSchedulerTrigger(uuid string, params param.UpdateSchedulerTriggerParam) (*view.SchedulerTriggerInventoryView, error) {
	var resp view.UpdateSchedulerTriggerEventView
	if err := cli.Put("v1/scheduler/triggers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteSchedulerTrigger deletes SchedulerTrigger
func (cli *ZSClient) DeleteSchedulerTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/triggers/{uuid}", uuid, string(deleteMode))
}
// CreateSchedulerTrigger creates SchedulerTrigger
func (cli *ZSClient) CreateSchedulerTrigger(params param.CreateSchedulerTriggerParam) (*view.SchedulerTriggerInventoryView, error) {
	var resp view.CreateSchedulerTriggerEventView
	if err := cli.Post("v1/scheduler/triggers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
