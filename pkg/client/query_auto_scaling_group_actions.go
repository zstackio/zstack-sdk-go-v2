// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAutoScalingGroup queries AutoScalingGroup list
func (cli *ZSClient) QueryAutoScalingGroup(params *param.QueryParam) ([]view.AutoScalingGroupInventoryView, error) {
	var resp []view.AutoScalingGroupInventoryView
	return resp, cli.List("v1/autoscaling/groups", params, &resp)
}
