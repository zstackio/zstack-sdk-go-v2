// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVpcUserVpnGatewayFromLocal queries VpcUserVpnGatewayFromLocal list
func (cli *ZSClient) QueryVpcUserVpnGatewayFromLocal(params *param.QueryParam) ([]view.VpcUserVpnGatewayInventoryView, error) {
	var resp []view.VpcUserVpnGatewayInventoryView
	return resp, cli.List("v1/hybrid/user-vpn", params, &resp)
}
