// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVpcIkeConfigFromLocal queries VpcIkeConfigFromLocal list
func (cli *ZSClient) QueryVpcIkeConfigFromLocal(params *param.QueryParam) ([]view.VpcVpnIkeConfigInventoryView, error) {
	var resp []view.VpcVpnIkeConfigInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection/ike", params, &resp)
}
