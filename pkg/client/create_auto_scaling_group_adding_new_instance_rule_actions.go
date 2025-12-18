// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAutoScalingGroupAddingNewInstanceRule creates AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) CreateAutoScalingGroupAddingNewInstanceRule(params param.CreateAutoScalingGroupAddingNewInstanceRuleParam) (*view.CreateAutoScalingRuleEventView, error) {
	resp := view.CreateAutoScalingRuleEventView{}
	if err := cli.Post("v1/autoscaling/rules/adding-new-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
