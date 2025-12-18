// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSchedulerJobGroup creates SchedulerJobGroup
func (cli *ZSClient) CreateSchedulerJobGroup(params param.CreateSchedulerJobGroupParam) (*view.CreateSchedulerJobGroupEventView, error) {
	resp := view.CreateSchedulerJobGroupEventView{}
	if err := cli.Post("v1/scheduler/jobgroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
