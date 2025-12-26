// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSchedulerJobsToSchedulerJobGroup adds SchedulerJobsToSchedulerJobGroup
func (cli *ZSClient) AddSchedulerJobsToSchedulerJobGroup(params param.AddSchedulerJobsToSchedulerJobGroupParam) (*view.AddSchedulerJobsToSchedulerJobGroupEventView, error) {
	resp := view.AddSchedulerJobsToSchedulerJobGroupEventView{}
	if err := cli.Post("v1/scheduler/jobgroups/{schedulerJobGroupUuid}/job", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
