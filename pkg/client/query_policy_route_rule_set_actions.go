// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPolicyRouteRuleSet queries PolicyRouteRuleSet list
func (cli *ZSClient) QueryPolicyRouteRuleSet(params *param.QueryParam) ([]view.PolicyRouteRuleSetInventoryView, error) {
	var resp []view.PolicyRouteRuleSetInventoryView
	return resp, cli.List("v1/policy-routes/rulesets", params, &resp)
}
