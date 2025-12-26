// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddHostToHostSchedulingRuleGroup adds HostToHostSchedulingRuleGroup
func (cli *ZSClient) AddHostToHostSchedulingRuleGroup(params param.AddHostToHostSchedulingRuleGroupParam) (*view.AddHostToHostSchedulingRuleGroupEventView, error) {
	resp := view.AddHostToHostSchedulingRuleGroupEventView{}
	if err := cli.Post("v1/hostSchedulingRuleGroup/{hostGroupUuid}/host/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
