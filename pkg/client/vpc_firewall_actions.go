// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVpcFirewall updates VpcFirewall
func (cli *ZSClient) UpdateVpcFirewall(ctx context.Context, uuid string, params param.UpdateVpcFirewallParam) (*view.VpcFirewallInventoryView, error) {
	resp := view.VpcFirewallInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vpcfirewalls", uuid, "", map[string]interface{}{
		"updateVpcFirewall": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateVpcFirewall creates VpcFirewall
func (cli *ZSClient) CreateVpcFirewall(ctx context.Context, params param.CreateVpcFirewallParam) (*view.VpcFirewallInventoryView, error) {
	resp := view.VpcFirewallInventoryView{}
	if err := cli.Post(ctx, "v1/vpcfirewalls", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVpcFirewall queries VpcFirewall list
func (cli *ZSClient) QueryVpcFirewall(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallInventoryView, error) {
	var resp []view.VpcFirewallInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls", params, &resp)
}

func (cli *ZSClient) GetVpcFirewall(ctx context.Context, uuid string) (*view.VpcFirewallInventoryView, error) {
	var resp view.VpcFirewallInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcFirewall Pagination
func (cli *ZSClient) PageVpcFirewall(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallInventoryView, int, error) {
	var vpcFirewalls []view.VpcFirewallInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls", params, &vpcFirewalls)
	return vpcFirewalls, total, err
}
