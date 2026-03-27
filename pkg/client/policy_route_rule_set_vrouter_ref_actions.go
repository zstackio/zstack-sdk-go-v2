// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPolicyRouteRuleSetVRouterRef queries PolicyRouteRuleSetVRouterRef list
func (cli *ZSClient) QueryPolicyRouteRuleSetVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteRuleSetVRouterRefInventoryView, error) {
	var resp []view.PolicyRouteRuleSetVRouterRefInventoryView
	return resp, cli.List(ctx, "v1/policy-routes/rulesets/vrouters/refs", params, &resp)
}

func (cli *ZSClient) GetPolicyRouteRuleSetVRouterRef(ctx context.Context, uuid string) (*view.PolicyRouteRuleSetVRouterRefInventoryView, error) {
	var resp view.PolicyRouteRuleSetVRouterRefInventoryView
	if err := cli.Get(ctx, "v1/policy-routes/rulesets/vrouters/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePolicyRouteRuleSetVRouterRef Pagination
func (cli *ZSClient) PagePolicyRouteRuleSetVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteRuleSetVRouterRefInventoryView, int, error) {
	var policyRouteRuleSetVRouterRefs []view.PolicyRouteRuleSetVRouterRefInventoryView
	total, err := cli.Page(ctx, "v1/policy-routes/rulesets/vrouters/refs", params, &policyRouteRuleSetVRouterRefs)
	return policyRouteRuleSetVRouterRefs, total, err
}
