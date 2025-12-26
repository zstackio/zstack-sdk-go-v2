// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAutoScalingGroupRemovalInstanceRule creates AutoScalingGroupRemovalInstanceRule
func (cli *ZSClient) CreateAutoScalingGroupRemovalInstanceRule(params param.CreateAutoScalingGroupRemovalInstanceRuleParam) (*view.CreateAutoScalingRuleEventView, error) {
	resp := view.CreateAutoScalingRuleEventView{}
	if err := cli.Post("v1/autoscaling/rules/removal-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
