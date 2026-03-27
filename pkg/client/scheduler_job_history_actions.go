// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySchedulerJobHistory queries SchedulerJobHistory list
func (cli *ZSClient) QuerySchedulerJobHistory(ctx context.Context, params *param.QueryParam) ([]view.SchedulerJobHistoryInventoryView, error) {
	var resp []view.SchedulerJobHistoryInventoryView
	return resp, cli.List(ctx, "v1/scheduler/job/history", params, &resp)
}

func (cli *ZSClient) GetSchedulerJobHistory(ctx context.Context, uuid string) (*view.SchedulerJobHistoryInventoryView, error) {
	var resp view.SchedulerJobHistoryInventoryView
	if err := cli.Get(ctx, "v1/scheduler/job/history", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSchedulerJobHistory Pagination
func (cli *ZSClient) PageSchedulerJobHistory(ctx context.Context, params *param.QueryParam) ([]view.SchedulerJobHistoryInventoryView, int, error) {
	var schedulerJobHistories []view.SchedulerJobHistoryInventoryView
	total, err := cli.Page(ctx, "v1/scheduler/job/history", params, &schedulerJobHistories)
	return schedulerJobHistories, total, err
}
