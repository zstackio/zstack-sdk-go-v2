// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSchedulerJob creates SchedulerJob
func (cli *ZSClient) CreateSchedulerJob(params param.CreateSchedulerJobParam) (*view.CreateSchedulerJobEventView, error) {
	resp := view.CreateSchedulerJobEventView{}
	if err := cli.Post("v1/scheduler/jobs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
