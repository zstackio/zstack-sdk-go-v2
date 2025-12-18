// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryL2PortGroupNetwork queries L2PortGroupNetwork list
func (cli *ZSClient) QueryL2PortGroupNetwork(params param.QueryParam) ([]view.L2PortGroupNetworkInventoryView, error) {
	var resp []view.L2PortGroupNetworkInventoryView
	return resp, cli.List("v1/l2-networks/port-group", &params, &resp)
}
