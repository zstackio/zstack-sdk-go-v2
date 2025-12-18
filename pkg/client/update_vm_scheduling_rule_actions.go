// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVmSchedulingRule updates VmSchedulingRule
func (cli *ZSClient) UpdateVmSchedulingRule(uuid string, params param.UpdateVmSchedulingRuleParam) (*view.UpdateVmSchedulingRuleEventView, error) {
	resp := view.UpdateVmSchedulingRuleEventView{}
	if err := cli.Put("v1/vmSchedulingRule/{uuid}/update", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
