// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPolicyRouteTableVRouterRef queries PolicyRouteTableVRouterRef list
func (cli *ZSClient) QueryPolicyRouteTableVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteTableVRouterRefInventoryView, error) {
	var resp []view.PolicyRouteTableVRouterRefInventoryView
	return resp, cli.List(ctx, "v1/policy-routes/tables/vrouters/refs", params, &resp)
}

func (cli *ZSClient) GetPolicyRouteTableVRouterRef(ctx context.Context, uuid string) (*view.PolicyRouteTableVRouterRefInventoryView, error) {
	var resp view.PolicyRouteTableVRouterRefInventoryView
	if err := cli.Get(ctx, "v1/policy-routes/tables/vrouters/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePolicyRouteTableVRouterRef Pagination
func (cli *ZSClient) PagePolicyRouteTableVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteTableVRouterRefInventoryView, int, error) {
	var policyRouteTableVRouterRefs []view.PolicyRouteTableVRouterRefInventoryView
	total, err := cli.Page(ctx, "v1/policy-routes/tables/vrouters/refs", params, &policyRouteTableVRouterRefs)
	return policyRouteTableVRouterRefs, total, err
}
