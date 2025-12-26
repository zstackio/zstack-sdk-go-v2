// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVpcIpSecConfigFromLocal queries VpcIpSecConfigFromLocal list
func (cli *ZSClient) QueryVpcIpSecConfigFromLocal(params *param.QueryParam) ([]view.VpcVpnIpSecConfigInventoryView, error) {
	var resp []view.VpcVpnIpSecConfigInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection/ipsec", params, &resp)
}
