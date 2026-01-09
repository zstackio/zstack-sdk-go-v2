// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVpcFirewallVRouterRef queries VpcFirewallVRouterRef list
func (cli *ZSClient) QueryVpcFirewallVRouterRef(params *param.QueryParam) ([]view.VpcFirewallVRouterRefInventoryView, error) {
	var resp []view.VpcFirewallVRouterRefInventoryView
	return resp, cli.List("v1/vpcfirewalls/vrouters/refs", params, &resp)
}

func (cli *ZSClient) GetVpcFirewallVRouterRef(uuid string) (*view.VpcFirewallVRouterRefInventoryView, error) {
	var resp view.VpcFirewallVRouterRefInventoryView
	if err := cli.Get("v1/vpcfirewalls/vrouters/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
