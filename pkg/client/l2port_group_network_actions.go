// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryL2PortGroupNetwork 查询L2PortGroupNetwork列表
func (cli *ZSClient) QueryL2PortGroupNetwork(params param.QueryParam) ([]view.QueryL2PortGroupNetworkView, error) {
	var resp []view.QueryL2PortGroupNetworkView
	return resp, cli.List("v1/l2-networks/port-group", &params, &resp)
}

