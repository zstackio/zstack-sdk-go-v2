// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcIpSecConfigFromLocal queries VpcIpSecConfigFromLocal list
func (cli *ZSClient) QueryVpcIpSecConfigFromLocal(params param.QueryParam) ([]view.VpcVpnIpSecConfigInventoryView, error) {
	var resp []view.VpcVpnIpSecConfigInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection/ipsec", &params, &resp)
}
