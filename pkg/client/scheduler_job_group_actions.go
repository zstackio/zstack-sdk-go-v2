// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

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
	if err := cli.PutWithRespKey("v1/scheduler/jobgroups", uuid, "", map[string]interface{}{
		"updateSchedulerJobGroup": params.Params,
	}, &resp); err != nil {
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

// AddSchedulerJobsToSchedulerJobGroup adds SchedulerJobsToSchedulerJobGroup
func (cli *ZSClient) AddSchedulerJobsToSchedulerJobGroup(schedulerJobGroupUuid string, params param.AddSchedulerJobsToSchedulerJobGroupParam) (*view.SchedulerJobGroupJobRefInventoryView, error) {
	resp := view.SchedulerJobGroupJobRefInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/scheduler/jobgroups/%s/job", schedulerJobGroupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveSchedulerJobGroupFromSchedulerTrigger removes SchedulerJobGroupFromSchedulerTrigger
func (cli *ZSClient) RemoveSchedulerJobGroupFromSchedulerTrigger(schedulerJobGroupUuid string, schedulerTriggerUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/scheduler/jobgroups", schedulerJobGroupUuid, fmt.Sprintf("scheduler/triggers/%s", schedulerTriggerUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// RemoveSchedulerJobsFromSchedulerJobGroup removes SchedulerJobsFromSchedulerJobGroup
func (cli *ZSClient) RemoveSchedulerJobsFromSchedulerJobGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobgroups", uuid, string(deleteMode))
}

// AddSchedulerJobGroupToSchedulerTrigger adds SchedulerJobGroupToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobGroupToSchedulerTrigger(schedulerJobGroupUuid string, schedulerTriggerUuid string, params param.AddSchedulerJobGroupToSchedulerTriggerParam) (*view.SchedulerJobGroupSchedulerTriggerRefInventoryView, error) {
	resp := view.SchedulerJobGroupSchedulerTriggerRefInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/scheduler/jobgroups/%s/scheduler/triggers/%s", schedulerJobGroupUuid, schedulerTriggerUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
