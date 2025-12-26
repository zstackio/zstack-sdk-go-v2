// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVpcVpnConnectionFromLocal queries VpcVpnConnectionFromLocal list
func (cli *ZSClient) QueryVpcVpnConnectionFromLocal(params *param.QueryParam) ([]view.VpcVpnConnectionInventoryView, error) {
	var resp []view.VpcVpnConnectionInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection", params, &resp)
}
