// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPolicyRouteRuleSetL3Ref queries PolicyRouteRuleSetL3Ref list
func (cli *ZSClient) QueryPolicyRouteRuleSetL3Ref(params *param.QueryParam) ([]view.PolicyRouteRuleSetL3RefInventoryView, error) {
	var resp []view.PolicyRouteRuleSetL3RefInventoryView
	return resp, cli.List("v1/policy-routes/rulesets/l3networdks/refs", params, &resp)
}
