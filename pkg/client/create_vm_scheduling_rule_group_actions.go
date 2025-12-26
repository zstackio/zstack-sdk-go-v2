// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmSchedulingRuleGroup creates VmSchedulingRuleGroup
func (cli *ZSClient) CreateVmSchedulingRuleGroup(params param.CreateVmSchedulingRuleGroupParam) (*view.CreateVmSchedulingRuleGroupEventView, error) {
	resp := view.CreateVmSchedulingRuleGroupEventView{}
	if err := cli.Post("v1/vmSchedulingRuleGroup", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
