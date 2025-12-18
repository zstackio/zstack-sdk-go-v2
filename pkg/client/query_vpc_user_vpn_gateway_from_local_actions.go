// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcUserVpnGatewayFromLocal queries VpcUserVpnGatewayFromLocal list
func (cli *ZSClient) QueryVpcUserVpnGatewayFromLocal(params param.QueryParam) ([]view.VpcUserVpnGatewayInventoryView, error) {
	var resp []view.VpcUserVpnGatewayInventoryView
	return resp, cli.List("v1/hybrid/user-vpn", &params, &resp)
}
