// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPolicyRouteRule queries PolicyRouteRule list
func (cli *ZSClient) QueryPolicyRouteRule(params *param.QueryParam) ([]view.PolicyRouteRuleInventoryView, error) {
	var resp []view.PolicyRouteRuleInventoryView
	return resp, cli.List("v1/policy-routes/rules", params, &resp)
}
