// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ValidateVmSchedulingRule operates on ValidateVmSchedulingRule
func (cli *ZSClient) ValidateVmSchedulingRule(uuid string, params param.ValidateVmSchedulingRuleParam) (*view.ValidateVmSchedulingRuleView, error) {
	resp := view.ValidateVmSchedulingRuleView{}
	if err := cli.Put("v1/validate/vmSchedulingRule", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
