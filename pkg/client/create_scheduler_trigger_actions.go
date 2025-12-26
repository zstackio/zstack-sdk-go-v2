// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSchedulerTrigger creates SchedulerTrigger
func (cli *ZSClient) CreateSchedulerTrigger(params param.CreateSchedulerTriggerParam) (*view.CreateSchedulerTriggerEventView, error) {
	resp := view.CreateSchedulerTriggerEventView{}
	if err := cli.Post("v1/scheduler/triggers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
