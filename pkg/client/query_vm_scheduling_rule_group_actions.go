// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVmSchedulingRuleGroup queries VmSchedulingRuleGroup list
func (cli *ZSClient) QueryVmSchedulingRuleGroup(params *param.QueryParam) ([]view.VmSchedulingRuleGroupInventoryView, error) {
	var resp []view.VmSchedulingRuleGroupInventoryView
	return resp, cli.List("v1/query/vm/schedulingRule/group", params, &resp)
}
