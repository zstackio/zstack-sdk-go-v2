// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSchedulerJobGroup updates SchedulerJobGroup
func (cli *ZSClient) UpdateSchedulerJobGroup(uuid string, params param.UpdateSchedulerJobGroupParam) (*view.UpdateSchedulerJobGroupEventView, error) {
	resp := view.UpdateSchedulerJobGroupEventView{}
	if err := cli.Put("v1/scheduler/jobgroups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
