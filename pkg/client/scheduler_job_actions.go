// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

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
	if err := cli.PutWithRespKey("v1/scheduler/jobs", uuid, "", map[string]interface{}{
		"updateSchedulerJob": params.Params,
	}, &resp); err != nil {
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

// RemoveSchedulerJobFromSchedulerTrigger removes SchedulerJobFromSchedulerTrigger
func (cli *ZSClient) RemoveSchedulerJobFromSchedulerTrigger(schedulerJobUuid string, schedulerTriggerUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/scheduler/jobs", schedulerJobUuid, fmt.Sprintf("scheduler/triggers/%s", schedulerTriggerUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetNoTriggerSchedulerJobs gets NoTriggerSchedulerJobs by uuid
func (cli *ZSClient) GetNoTriggerSchedulerJobs() (*view.SchedulerJobInventoryView, error) {
	var resp view.SchedulerJobInventoryView
	if err := cli.GetWithRespKey("v1/scheduler/jobs/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSchedulerJobToSchedulerTrigger adds SchedulerJobToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobToSchedulerTrigger(schedulerJobUuid string, schedulerTriggerUuid string, params param.AddSchedulerJobToSchedulerTriggerParam) (*view.SchedulerJobSchedulerTriggerInventoryView, error) {
	resp := view.SchedulerJobSchedulerTriggerInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/scheduler/jobs/%s/scheduler/triggers/%s", schedulerJobUuid, schedulerTriggerUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSchedulerState changes SchedulerState
func (cli *ZSClient) ChangeSchedulerState(uuid string, params param.ChangeSchedulerStateParam) (*view.SchedulerJobInventoryView, error) {
	resp := view.SchedulerJobInventoryView{}
	if err := cli.PutWithRespKey("v1/schedulers", uuid, "", map[string]interface{}{
		"changeSchedulerState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
