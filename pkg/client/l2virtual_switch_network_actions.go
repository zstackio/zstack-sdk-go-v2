// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryL2VirtualSwitchNetwork 查询L2VirtualSwitchNetwork列表
func (cli *ZSClient) QueryL2VirtualSwitchNetwork(params param.QueryParam) ([]view.QueryL2VirtualSwitchNetworkView, error) {
	var resp []view.QueryL2VirtualSwitchNetworkView
	return resp, cli.List("v1/l2-networks/virtual-switch", &params, &resp)
}

