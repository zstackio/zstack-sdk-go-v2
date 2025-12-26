// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSchedulerJobToSchedulerTrigger adds SchedulerJobToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobToSchedulerTrigger(params param.AddSchedulerJobToSchedulerTriggerParam) (*view.AddSchedulerJobToSchedulerTriggerEventView, error) {
	resp := view.AddSchedulerJobToSchedulerTriggerEventView{}
	if err := cli.Post("v1/scheduler/jobs/{schedulerJobUuid}/scheduler/triggers/{schedulerTriggerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
