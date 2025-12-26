// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ListVmSchedulingRulesFromExecuteState operates on ListVmSchedulingRulesFromExecuteState
func (cli *ZSClient) ListVmSchedulingRulesFromExecuteState(params param.ListVmSchedulingRulesFromExecuteStateParam) (*view.ListVmSchedulingRulesFromExecuteStateView, error) {
	resp := view.ListVmSchedulingRulesFromExecuteStateView{}
	if err := cli.Post("v1/list/vmSchedulingRules/from/conflict/state", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
