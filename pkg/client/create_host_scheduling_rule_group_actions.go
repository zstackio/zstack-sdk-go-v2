// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateHostSchedulingRuleGroup creates HostSchedulingRuleGroup
func (cli *ZSClient) CreateHostSchedulingRuleGroup(params param.CreateHostSchedulingRuleGroupParam) (*view.CreateHostSchedulingRuleGroupEventView, error) {
	resp := view.CreateHostSchedulingRuleGroupEventView{}
	if err := cli.Post("v1/hostSchedulingRuleGroup", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
