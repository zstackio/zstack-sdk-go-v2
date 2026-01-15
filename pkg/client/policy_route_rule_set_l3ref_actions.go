// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPolicyRouteRuleSetL3Ref queries PolicyRouteRuleSetL3Ref list
func (cli *ZSClient) QueryPolicyRouteRuleSetL3Ref(params *param.QueryParam) ([]view.PolicyRouteRuleSetL3RefInventoryView, error) {
	var resp []view.PolicyRouteRuleSetL3RefInventoryView
	return resp, cli.List("v1/policy-routes/rulesets/l3networdks/refs", params, &resp)
}

// PagePolicyRouteRuleSetL3Ref Pagination
func (cli *ZSClient) PagePolicyRouteRuleSetL3Ref(params *param.QueryParam) ([]view.PolicyRouteRuleSetL3RefInventoryView, int, error) {
	var policyRouteRuleSetL3Refs []view.PolicyRouteRuleSetL3RefInventoryView
	total, err := cli.Page("v1/policy-routes/rulesets/l3networdks/refs", params, &policyRouteRuleSetL3Refs)
	return policyRouteRuleSetL3Refs, total, err
}
