// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcVpnConnectionFromLocal 查询VpcVpnConnectionFromLocal列表
func (cli *ZSClient) QueryVpcVpnConnectionFromLocal(params param.QueryParam) ([]view.QueryVpcVpnConnectionFromLocalView, error) {
	var resp []view.QueryVpcVpnConnectionFromLocalView
	return resp, cli.List("v1/hybrid/vpn-connection", &params, &resp)
}

