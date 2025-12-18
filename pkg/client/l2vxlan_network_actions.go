// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryL2VxlanNetwork 查询L2VxlanNetwork列表
func (cli *ZSClient) QueryL2VxlanNetwork(params param.QueryParam) ([]view.QueryL2VxlanNetworkView, error) {
	var resp []view.QueryL2VxlanNetworkView
	return resp, cli.List("v1/l2-networks/vxlan", &params, &resp)
}

