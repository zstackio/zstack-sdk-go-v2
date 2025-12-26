// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVmSchedulingRuleGroup updates VmSchedulingRuleGroup
func (cli *ZSClient) UpdateVmSchedulingRuleGroup(uuid string, params param.UpdateVmSchedulingRuleGroupParam) (*view.UpdateVmSchedulingRuleGroupEventView, error) {
	resp := view.UpdateVmSchedulingRuleGroupEventView{}
	if err := cli.Put("v1/vmSchedulingRuleGroup/{uuid}/update", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
