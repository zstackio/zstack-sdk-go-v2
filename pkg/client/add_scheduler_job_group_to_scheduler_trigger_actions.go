// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSchedulerJobGroupToSchedulerTrigger adds SchedulerJobGroupToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobGroupToSchedulerTrigger(params param.AddSchedulerJobGroupToSchedulerTriggerParam) (*view.AddSchedulerJobGroupToSchedulerTriggerEventView, error) {
	resp := view.AddSchedulerJobGroupToSchedulerTriggerEventView{}
	if err := cli.Post("v1/scheduler/jobgroups/{schedulerJobGroupUuid}/scheduler/triggers/{schedulerTriggerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
