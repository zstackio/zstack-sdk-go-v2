// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAutoScalingGroupRemovalInstanceRule updates AutoScalingGroupRemovalInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupRemovalInstanceRule(uuid string, params param.UpdateAutoScalingGroupRemovalInstanceRuleParam) (*view.UpdateAutoScalingRuleEventView, error) {
	resp := view.UpdateAutoScalingRuleEventView{}
	if err := cli.Put("v1/autoscaling/rules/removal-instance/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
