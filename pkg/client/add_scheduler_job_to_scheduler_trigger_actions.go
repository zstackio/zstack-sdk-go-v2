// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSchedulerJobToSchedulerTrigger 操作AddSchedulerJobToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobToSchedulerTrigger(params param.AddSchedulerJobToSchedulerTriggerParam) (*view.AddSchedulerJobToSchedulerTriggerEventView, error) {
	resp := view.AddSchedulerJobToSchedulerTriggerEventView{}
	if err := cli.Post("v1/scheduler/jobs/{schedulerJobUuid}/scheduler/triggers/{schedulerTriggerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

