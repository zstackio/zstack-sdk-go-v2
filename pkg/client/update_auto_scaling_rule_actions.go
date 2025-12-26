// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAutoScalingRule updates AutoScalingRule
func (cli *ZSClient) UpdateAutoScalingRule(uuid string, params param.UpdateAutoScalingRuleParam) (*view.UpdateAutoScalingRuleEventView, error) {
	resp := view.UpdateAutoScalingRuleEventView{}
	if err := cli.Put("v1/autoscaling/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
