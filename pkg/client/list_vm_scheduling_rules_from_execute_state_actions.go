// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ListVmSchedulingRulesFromExecuteState 操作ListVmSchedulingRulesFromExecuteState
func (cli *ZSClient) ListVmSchedulingRulesFromExecuteState(params param.ListVmSchedulingRulesFromExecuteStateParam) (*view.ListVmSchedulingRulesFromExecuteStateView, error) {
	resp := view.ListVmSchedulingRulesFromExecuteStateView{}
	if err := cli.Post("v1/list/vmSchedulingRules/from/conflict/state", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

