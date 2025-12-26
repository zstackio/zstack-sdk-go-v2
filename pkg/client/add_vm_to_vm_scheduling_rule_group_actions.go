// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddVmToVmSchedulingRuleGroup adds VmToVmSchedulingRuleGroup
func (cli *ZSClient) AddVmToVmSchedulingRuleGroup(params param.AddVmToVmSchedulingRuleGroupParam) (*view.AddVmToVmSchedulingRuleGroupEventView, error) {
	resp := view.AddVmToVmSchedulingRuleGroupEventView{}
	if err := cli.Post("v1/vmSchedulingRuleGroup/{vmGroupUuid}/vmInstance/{vmUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
