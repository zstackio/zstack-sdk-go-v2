// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVmSchedulingRuleState changes VmSchedulingRuleState
func (cli *ZSClient) ChangeVmSchedulingRuleState(uuid string, params param.ChangeVmSchedulingRuleStateParam) (*view.ChangeVmSchedulingRuleStateEventView, error) {
	resp := view.ChangeVmSchedulingRuleStateEventView{}
	if err := cli.Put("v1/vmSchedulingRule/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
