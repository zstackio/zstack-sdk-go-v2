// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcIkeConfigFromLocal 查询VpcIkeConfigFromLocal列表
func (cli *ZSClient) QueryVpcIkeConfigFromLocal(params param.QueryParam) ([]view.QueryVpcIkeConfigFromLocalView, error) {
	var resp []view.QueryVpcIkeConfigFromLocalView
	return resp, cli.List("v1/hybrid/vpn-connection/ike", &params, &resp)
}

