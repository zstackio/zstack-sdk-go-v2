// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSchedulerTrigger creates SchedulerTrigger
func (cli *ZSClient) CreateSchedulerTrigger(params param.CreateSchedulerTriggerParam) (*view.CreateSchedulerTriggerEventView, error) {
	resp := view.CreateSchedulerTriggerEventView{}
	if err := cli.Post("v1/scheduler/triggers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
