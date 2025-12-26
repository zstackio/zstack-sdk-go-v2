// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAutoScalingRule queries AutoScalingRule list
func (cli *ZSClient) QueryAutoScalingRule(params *param.QueryParam) ([]view.AutoScalingRuleInventoryView, error) {
	var resp []view.AutoScalingRuleInventoryView
	return resp, cli.List("v1/autoscaling/groups/rules", params, &resp)
}
