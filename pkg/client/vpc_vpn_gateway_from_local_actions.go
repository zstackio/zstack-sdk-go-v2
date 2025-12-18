// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcVpnGatewayFromLocal 查询VpcVpnGatewayFromLocal列表
func (cli *ZSClient) QueryVpcVpnGatewayFromLocal(params param.QueryParam) ([]view.QueryVpcVpnGatewayFromLocalView, error) {
	var resp []view.QueryVpcVpnGatewayFromLocalView
	return resp, cli.List("v1/hybrid/vpc-vpn", &params, &resp)
}

