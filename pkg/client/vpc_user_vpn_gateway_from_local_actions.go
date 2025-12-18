// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcUserVpnGatewayFromLocal 查询VpcUserVpnGatewayFromLocal列表
func (cli *ZSClient) QueryVpcUserVpnGatewayFromLocal(params param.QueryParam) ([]view.QueryVpcUserVpnGatewayFromLocalView, error) {
	var resp []view.QueryVpcUserVpnGatewayFromLocalView
	return resp, cli.List("v1/hybrid/user-vpn", &params, &resp)
}

