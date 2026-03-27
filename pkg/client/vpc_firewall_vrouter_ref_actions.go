// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVpcFirewallVRouterRef queries VpcFirewallVRouterRef list
func (cli *ZSClient) QueryVpcFirewallVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallVRouterRefInventoryView, error) {
	var resp []view.VpcFirewallVRouterRefInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/vrouters/refs", params, &resp)
}

func (cli *ZSClient) GetVpcFirewallVRouterRef(ctx context.Context, uuid string) (*view.VpcFirewallVRouterRefInventoryView, error) {
	var resp view.VpcFirewallVRouterRefInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls/vrouters/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcFirewallVRouterRef Pagination
func (cli *ZSClient) PageVpcFirewallVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallVRouterRefInventoryView, int, error) {
	var vpcFirewallVRouterRefs []view.VpcFirewallVRouterRefInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/vrouters/refs", params, &vpcFirewallVRouterRefs)
	return vpcFirewallVRouterRefs, total, err
}
