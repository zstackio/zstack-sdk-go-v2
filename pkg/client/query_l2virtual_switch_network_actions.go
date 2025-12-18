// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryL2VirtualSwitchNetwork queries L2VirtualSwitchNetwork list
func (cli *ZSClient) QueryL2VirtualSwitchNetwork(params param.QueryParam) ([]view.L2VirtualSwitchNetworkInventoryView, error) {
	var resp []view.L2VirtualSwitchNetworkInventoryView
	return resp, cli.List("v1/l2-networks/virtual-switch", &params, &resp)
}
