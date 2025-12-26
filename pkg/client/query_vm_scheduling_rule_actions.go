// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVmSchedulingRule queries VmSchedulingRule list
func (cli *ZSClient) QueryVmSchedulingRule(params *param.QueryParam) ([]view.VmSchedulingRuleInventoryView, error) {
	var resp []view.VmSchedulingRuleInventoryView
	return resp, cli.List("v1/query/vm/schedulingRule", params, &resp)
}
