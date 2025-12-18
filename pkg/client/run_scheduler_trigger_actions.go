// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RunSchedulerTrigger 操作RunSchedulerTrigger
func (cli *ZSClient) RunSchedulerTrigger(uuid string, params param.RunSchedulerTriggerParam) (*view.RunSchedulerTriggerEventView, error) {
	resp := view.RunSchedulerTriggerEventView{}
	if err := cli.Put("v1/scheduler/triggers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

