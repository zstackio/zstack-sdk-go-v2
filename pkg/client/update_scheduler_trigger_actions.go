// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSchedulerTrigger updates SchedulerTrigger
func (cli *ZSClient) UpdateSchedulerTrigger(uuid string, params param.UpdateSchedulerTriggerParam) (*view.UpdateSchedulerTriggerEventView, error) {
	resp := view.UpdateSchedulerTriggerEventView{}
	if err := cli.Put("v1/scheduler/triggers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
