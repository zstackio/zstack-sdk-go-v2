// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVpcVpnGatewayFromLocal queries VpcVpnGatewayFromLocal list
func (cli *ZSClient) QueryVpcVpnGatewayFromLocal(params *param.QueryParam) ([]view.VpcVpnGatewayInventoryView, error) {
	var resp []view.VpcVpnGatewayInventoryView
	return resp, cli.List("v1/hybrid/vpc-vpn", params, &resp)
}
