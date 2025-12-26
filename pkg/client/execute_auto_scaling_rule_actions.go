// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ExecuteAutoScalingRule operates on ExecuteAutoScalingRule
func (cli *ZSClient) ExecuteAutoScalingRule(uuid string, params param.ExecuteAutoScalingRuleParam) (*view.ExecuteAutoScalingRuleEventView, error) {
	resp := view.ExecuteAutoScalingRuleEventView{}
	if err := cli.Put("v1/autoscaling/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
