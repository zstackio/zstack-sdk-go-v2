// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySchedulerTrigger queries SchedulerTrigger list
func (cli *ZSClient) QuerySchedulerTrigger(params *param.QueryParam) ([]view.SchedulerTriggerInventoryView, error) {
	var resp []view.SchedulerTriggerInventoryView
	return resp, cli.List("v1/scheduler/triggers", params, &resp)
}
