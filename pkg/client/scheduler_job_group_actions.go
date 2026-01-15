// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSchedulerJobGroup creates SchedulerJobGroup
func (cli *ZSClient) CreateSchedulerJobGroup(params param.CreateSchedulerJobGroupParam) (*view.SchedulerJobGroupInventoryView, error) {
	resp := view.SchedulerJobGroupInventoryView{}
	if err := cli.Post("v1/scheduler/jobgroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSchedulerJobGroup deletes SchedulerJobGroup
func (cli *ZSClient) DeleteSchedulerJobGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobgroups", uuid, string(deleteMode))
}
// UpdateSchedulerJobGroup updates SchedulerJobGroup
func (cli *ZSClient) UpdateSchedulerJobGroup(uuid string, params param.UpdateSchedulerJobGroupParam) (*view.SchedulerJobGroupInventoryView, error) {
	resp := view.SchedulerJobGroupInventoryView{}
	if err := cli.Put("v1/scheduler/jobgroups", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySchedulerJobGroup queries SchedulerJobGroup list
func (cli *ZSClient) QuerySchedulerJobGroup(params *param.QueryParam) ([]view.SchedulerJobGroupInventoryView, error) {
	var resp []view.SchedulerJobGroupInventoryView
	return resp, cli.List("v1/scheduler/jobgroups", params, &resp)
}

func (cli *ZSClient) GetSchedulerJobGroup(uuid string) (*view.SchedulerJobGroupInventoryView, error) {
	var resp view.SchedulerJobGroupInventoryView
	if err := cli.Get("v1/scheduler/jobgroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSchedulerJobGroup Pagination
func (cli *ZSClient) PageSchedulerJobGroup(params *param.QueryParam) ([]view.SchedulerJobGroupInventoryView, int, error) {
	var schedulerJobGroups []view.SchedulerJobGroupInventoryView
	total, err := cli.Page("v1/scheduler/jobgroups", params, &schedulerJobGroups)
	return schedulerJobGroups, total, err
}
