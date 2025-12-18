// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAutoScalingRule queries AutoScalingRule list
func (cli *ZSClient) QueryAutoScalingRule(params param.QueryParam) ([]view.AutoScalingRuleInventoryView, error) {
	var resp []view.AutoScalingRuleInventoryView
	return resp, cli.List("v1/autoscaling/groups/rules", &params, &resp)
}
