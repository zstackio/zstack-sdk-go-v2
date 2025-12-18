// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcVpnConnectionFromLocal queries VpcVpnConnectionFromLocal list
func (cli *ZSClient) QueryVpcVpnConnectionFromLocal(params param.QueryParam) ([]view.VpcVpnConnectionInventoryView, error) {
	var resp []view.VpcVpnConnectionInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection", &params, &resp)
}
