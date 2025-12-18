// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcIkeConfigFromLocal queries VpcIkeConfigFromLocal list
func (cli *ZSClient) QueryVpcIkeConfigFromLocal(params param.QueryParam) ([]view.VpcVpnIkeConfigInventoryView, error) {
	var resp []view.VpcVpnIkeConfigInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection/ike", &params, &resp)
}
