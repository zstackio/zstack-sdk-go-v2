// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAutoScalingGroupAddingNewInstanceRule updates AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupAddingNewInstanceRule(uuid string, params param.UpdateAutoScalingGroupAddingNewInstanceRuleParam) (*view.UpdateAutoScalingRuleEventView, error) {
	resp := view.UpdateAutoScalingRuleEventView{}
	if err := cli.Put("v1/autoscaling/rules/adding-new-instance/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
