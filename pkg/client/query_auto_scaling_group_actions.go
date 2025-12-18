// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAutoScalingGroup queries AutoScalingGroup list
func (cli *ZSClient) QueryAutoScalingGroup(params param.QueryParam) ([]view.AutoScalingGroupInventoryView, error) {
	var resp []view.AutoScalingGroupInventoryView
	return resp, cli.List("v1/autoscaling/groups", &params, &resp)
}
