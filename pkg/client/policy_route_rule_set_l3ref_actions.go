// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPolicyRouteRuleSetL3Ref queries PolicyRouteRuleSetL3Ref list
func (cli *ZSClient) QueryPolicyRouteRuleSetL3Ref(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteRuleSetL3RefInventoryView, error) {
	var resp []view.PolicyRouteRuleSetL3RefInventoryView
	return resp, cli.List(ctx, "v1/policy-routes/rulesets/l3networdks/refs", params, &resp)
}

func (cli *ZSClient) GetPolicyRouteRuleSetL3Ref(ctx context.Context, uuid string) (*view.PolicyRouteRuleSetL3RefInventoryView, error) {
	var resp view.PolicyRouteRuleSetL3RefInventoryView
	if err := cli.Get(ctx, "v1/policy-routes/rulesets/l3networdks/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePolicyRouteRuleSetL3Ref Pagination
func (cli *ZSClient) PagePolicyRouteRuleSetL3Ref(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteRuleSetL3RefInventoryView, int, error) {
	var policyRouteRuleSetL3Refs []view.PolicyRouteRuleSetL3RefInventoryView
	total, err := cli.Page(ctx, "v1/policy-routes/rulesets/l3networdks/refs", params, &policyRouteRuleSetL3Refs)
	return policyRouteRuleSetL3Refs, total, err
}
