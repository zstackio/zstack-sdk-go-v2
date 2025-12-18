// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAutoScalingRuleTrigger queries AutoScalingRuleTrigger list
func (cli *ZSClient) QueryAutoScalingRuleTrigger(params param.QueryParam) ([]view.AutoScalingRuleTriggerInventoryView, error) {
	var resp []view.AutoScalingRuleTriggerInventoryView
	return resp, cli.List("v1/autoscaling/groups/rules/trigger", &params, &resp)
}
