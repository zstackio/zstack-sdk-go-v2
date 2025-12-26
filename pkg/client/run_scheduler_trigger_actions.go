// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RunSchedulerTrigger operates on RunSchedulerTrigger
func (cli *ZSClient) RunSchedulerTrigger(uuid string, params param.RunSchedulerTriggerParam) (*view.RunSchedulerTriggerEventView, error) {
	resp := view.RunSchedulerTriggerEventView{}
	if err := cli.Put("v1/scheduler/triggers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
