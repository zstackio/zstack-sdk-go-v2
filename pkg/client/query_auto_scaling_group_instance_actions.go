// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAutoScalingGroupInstance queries AutoScalingGroupInstance list
func (cli *ZSClient) QueryAutoScalingGroupInstance(params *param.QueryParam) ([]view.AutoScalingGroupInstanceInventoryView, error) {
	var resp []view.AutoScalingGroupInstanceInventoryView
	return resp, cli.List("v1/autoscaling/groups/instances", params, &resp)
}
