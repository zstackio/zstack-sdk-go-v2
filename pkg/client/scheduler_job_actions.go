// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSchedulerJob creates SchedulerJob
func (cli *ZSClient) CreateSchedulerJob(params param.CreateSchedulerJobParam) (*view.SchedulerJobInventoryView, error) {
	var resp view.CreateSchedulerJobEventView
	if err := cli.Post("v1/scheduler/jobs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateSchedulerJob updates SchedulerJob
func (cli *ZSClient) UpdateSchedulerJob(uuid string, params param.UpdateSchedulerJobParam) (*view.SchedulerJobInventoryView, error) {
	var resp view.UpdateSchedulerJobEventView
	if err := cli.Put("v1/scheduler/jobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteSchedulerJob deletes SchedulerJob
func (cli *ZSClient) DeleteSchedulerJob(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobs/{uuid}", uuid, string(deleteMode))
}
// QuerySchedulerJob queries SchedulerJob list
func (cli *ZSClient) QuerySchedulerJob(params *param.QueryParam) ([]view.SchedulerJobInventoryView, error) {
	var resp []view.SchedulerJobInventoryView
	return resp, cli.List("v1/scheduler/jobs", params, &resp)
}
