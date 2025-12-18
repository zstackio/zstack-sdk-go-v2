// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySchedulerTrigger queries SchedulerTrigger list
func (cli *ZSClient) QuerySchedulerTrigger(params param.QueryParam) ([]view.SchedulerTriggerInventoryView, error) {
	var resp []view.SchedulerTriggerInventoryView
	return resp, cli.List("v1/scheduler/triggers", &params, &resp)
}
