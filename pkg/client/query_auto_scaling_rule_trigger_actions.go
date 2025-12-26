// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAutoScalingRuleTrigger queries AutoScalingRuleTrigger list
func (cli *ZSClient) QueryAutoScalingRuleTrigger(params *param.QueryParam) ([]view.AutoScalingRuleTriggerInventoryView, error) {
	var resp []view.AutoScalingRuleTriggerInventoryView
	return resp, cli.List("v1/autoscaling/groups/rules/trigger", params, &resp)
}
