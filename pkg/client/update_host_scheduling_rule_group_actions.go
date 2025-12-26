// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateHostSchedulingRuleGroup updates HostSchedulingRuleGroup
func (cli *ZSClient) UpdateHostSchedulingRuleGroup(uuid string, params param.UpdateHostSchedulingRuleGroupParam) (*view.UpdateHostSchedulingRuleGroupEventView, error) {
	resp := view.UpdateHostSchedulingRuleGroupEventView{}
	if err := cli.Put("v1/hostSchedulingRuleGroup/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
