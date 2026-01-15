// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSchedulerJob creates SchedulerJob
func (cli *ZSClient) CreateSchedulerJob(params param.CreateSchedulerJobParam) (*view.SchedulerJobInventoryView, error) {
	resp := view.SchedulerJobInventoryView{}
	if err := cli.Post("v1/scheduler/jobs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSchedulerJob updates SchedulerJob
func (cli *ZSClient) UpdateSchedulerJob(uuid string, params param.UpdateSchedulerJobParam) (*view.SchedulerJobInventoryView, error) {
	resp := view.SchedulerJobInventoryView{}
	if err := cli.Put("v1/scheduler/jobs", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSchedulerJob deletes SchedulerJob
func (cli *ZSClient) DeleteSchedulerJob(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobs", uuid, string(deleteMode))
}
// QuerySchedulerJob queries SchedulerJob list
func (cli *ZSClient) QuerySchedulerJob(params *param.QueryParam) ([]view.SchedulerJobInventoryView, error) {
	var resp []view.SchedulerJobInventoryView
	return resp, cli.List("v1/scheduler/jobs", params, &resp)
}

func (cli *ZSClient) GetSchedulerJob(uuid string) (*view.SchedulerJobInventoryView, error) {
	var resp view.SchedulerJobInventoryView
	if err := cli.Get("v1/scheduler/jobs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSchedulerJob Pagination
func (cli *ZSClient) PageSchedulerJob(params *param.QueryParam) ([]view.SchedulerJobInventoryView, int, error) {
	var schedulerJobs []view.SchedulerJobInventoryView
	total, err := cli.Page("v1/scheduler/jobs", params, &schedulerJobs)
	return schedulerJobs, total, err
}
