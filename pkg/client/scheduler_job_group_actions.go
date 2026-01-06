// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSchedulerJobGroup creates SchedulerJobGroup
func (cli *ZSClient) CreateSchedulerJobGroup(params param.CreateSchedulerJobGroupParam) (*view.SchedulerJobGroupInventoryView, error) {
	var resp view.CreateSchedulerJobGroupEventView
	if err := cli.Post("v1/scheduler/jobgroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteSchedulerJobGroup deletes SchedulerJobGroup
func (cli *ZSClient) DeleteSchedulerJobGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobgroups/{uuid}", uuid, string(deleteMode))
}
// UpdateSchedulerJobGroup updates SchedulerJobGroup
func (cli *ZSClient) UpdateSchedulerJobGroup(uuid string, params param.UpdateSchedulerJobGroupParam) (*view.SchedulerJobGroupInventoryView, error) {
	var resp view.UpdateSchedulerJobGroupEventView
	if err := cli.Put("v1/scheduler/jobgroups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySchedulerJobGroup queries SchedulerJobGroup list
func (cli *ZSClient) QuerySchedulerJobGroup(params *param.QueryParam) ([]view.SchedulerJobGroupInventoryView, error) {
	var resp []view.SchedulerJobGroupInventoryView
	return resp, cli.List("v1/scheduler/jobgroups", params, &resp)
}
