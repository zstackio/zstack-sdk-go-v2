// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAutoScalingGroupAddingNewInstanceRule creates AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) CreateAutoScalingGroupAddingNewInstanceRule(params param.CreateAutoScalingGroupAddingNewInstanceRuleParam) (*view.CreateAutoScalingRuleEventView, error) {
	resp := view.CreateAutoScalingRuleEventView{}
	if err := cli.Post("v1/autoscaling/rules/adding-new-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
