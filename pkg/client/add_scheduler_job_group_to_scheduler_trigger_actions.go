// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSchedulerJobGroupToSchedulerTrigger 操作AddSchedulerJobGroupToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobGroupToSchedulerTrigger(params param.AddSchedulerJobGroupToSchedulerTriggerParam) (*view.AddSchedulerJobGroupToSchedulerTriggerEventView, error) {
	resp := view.AddSchedulerJobGroupToSchedulerTriggerEventView{}
	if err := cli.Post("v1/scheduler/jobgroups/{schedulerJobGroupUuid}/scheduler/triggers/{schedulerTriggerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

