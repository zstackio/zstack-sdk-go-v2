// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPolicyRouteRuleSetVRouterRef queries PolicyRouteRuleSetVRouterRef list
func (cli *ZSClient) QueryPolicyRouteRuleSetVRouterRef(params *param.QueryParam) ([]view.PolicyRouteRuleSetVRouterRefInventoryView, error) {
	var resp []view.PolicyRouteRuleSetVRouterRefInventoryView
	return resp, cli.List("v1/policy-routes/rulesets/vrouters/refs", params, &resp)
}
