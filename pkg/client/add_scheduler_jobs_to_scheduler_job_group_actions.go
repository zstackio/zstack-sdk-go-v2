// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSchedulerJobsToSchedulerJobGroup adds SchedulerJobsToSchedulerJobGroup
func (cli *ZSClient) AddSchedulerJobsToSchedulerJobGroup(params param.AddSchedulerJobsToSchedulerJobGroupParam) (*view.AddSchedulerJobsToSchedulerJobGroupEventView, error) {
	resp := view.AddSchedulerJobsToSchedulerJobGroupEventView{}
	if err := cli.Post("v1/scheduler/jobgroups/{schedulerJobGroupUuid}/job", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
