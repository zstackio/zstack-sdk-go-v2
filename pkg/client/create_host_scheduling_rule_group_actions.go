// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateHostSchedulingRuleGroup creates HostSchedulingRuleGroup
func (cli *ZSClient) CreateHostSchedulingRuleGroup(params param.CreateHostSchedulingRuleGroupParam) (*view.CreateHostSchedulingRuleGroupEventView, error) {
	resp := view.CreateHostSchedulingRuleGroupEventView{}
	if err := cli.Post("v1/hostSchedulingRuleGroup", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
