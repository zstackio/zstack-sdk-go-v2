// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVpcFirewall updates VpcFirewall
func (cli *ZSClient) UpdateVpcFirewall(uuid string, params param.UpdateVpcFirewallParam) (*view.VpcFirewallInventoryView, error) {
	var resp view.UpdateVpcFirewallEventView
	if err := cli.Put("v1/vpcfirewalls/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateVpcFirewall creates VpcFirewall
func (cli *ZSClient) CreateVpcFirewall(params param.CreateVpcFirewallParam) (*view.VpcFirewallInventoryView, error) {
	var resp view.CreateVpcFirewallEventView
	if err := cli.Post("v1/vpcfirewalls", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVpcFirewall queries VpcFirewall list
func (cli *ZSClient) QueryVpcFirewall(params *param.QueryParam) ([]view.VpcFirewallInventoryView, error) {
	var resp []view.VpcFirewallInventoryView
	return resp, cli.List("v1/vpcfirewalls", params, &resp)
}
