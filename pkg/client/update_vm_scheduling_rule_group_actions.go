// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVmSchedulingRuleGroup updates VmSchedulingRuleGroup
func (cli *ZSClient) UpdateVmSchedulingRuleGroup(uuid string, params param.UpdateVmSchedulingRuleGroupParam) (*view.UpdateVmSchedulingRuleGroupEventView, error) {
	resp := view.UpdateVmSchedulingRuleGroupEventView{}
	if err := cli.Put("v1/vmSchedulingRuleGroup/{uuid}/update", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
