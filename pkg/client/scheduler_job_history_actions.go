// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySchedulerJobHistory queries SchedulerJobHistory list
func (cli *ZSClient) QuerySchedulerJobHistory(params *param.QueryParam) ([]view.SchedulerJobHistoryInventoryView, error) {
	var resp []view.SchedulerJobHistoryInventoryView
	return resp, cli.List("v1/scheduler/job/history", params, &resp)
}

func (cli *ZSClient) GetSchedulerJobHistory(uuid string) (*view.SchedulerJobHistoryInventoryView, error) {
	var resp view.SchedulerJobHistoryInventoryView
	if err := cli.Get("v1/scheduler/job/history", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSchedulerJobHistory Pagination
func (cli *ZSClient) PageSchedulerJobHistory(params *param.QueryParam) ([]view.SchedulerJobHistoryInventoryView, int, error) {
	var schedulerJobHistories []view.SchedulerJobHistoryInventoryView
	total, err := cli.Page("v1/scheduler/job/history", params, &schedulerJobHistories)
	return schedulerJobHistories, total, err
}
