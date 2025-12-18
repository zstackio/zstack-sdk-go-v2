// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmSchedulingRulesExecuteState gets VmSchedulingRulesExecuteState by uuid
func (cli *ZSClient) GetVmSchedulingRulesExecuteState(uuid string) (*view.GetVmSchedulingRulesExecuteStateView, error) {
	var resp view.GetVmSchedulingRulesExecuteStateView
	if err := cli.Get("v1/get/vmSchedulingRules/conflict/state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
