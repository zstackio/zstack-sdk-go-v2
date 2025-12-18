// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAutoScalingGroupAddingNewInstanceRule 更新AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupAddingNewInstanceRule(uuid string, params param.UpdateAutoScalingGroupAddingNewInstanceRuleParam) (*view.UpdateAutoScalingRuleEventView, error) {
	resp := view.UpdateAutoScalingRuleEventView{}
	if err := cli.Put("v1/autoscaling/rules/adding-new-instance/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

