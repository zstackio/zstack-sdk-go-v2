// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcIpSecConfigFromLocal 查询VpcIpSecConfigFromLocal列表
func (cli *ZSClient) QueryVpcIpSecConfigFromLocal(params param.QueryParam) ([]view.QueryVpcIpSecConfigFromLocalView, error) {
	var resp []view.QueryVpcIpSecConfigFromLocalView
	return resp, cli.List("v1/hybrid/vpn-connection/ipsec", &params, &resp)
}

