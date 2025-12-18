// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryL2VlanNetwork queries L2VlanNetwork list
func (cli *ZSClient) QueryL2VlanNetwork(params param.QueryParam) ([]view.L2VlanNetworkInventoryView, error) {
	var resp []view.L2VlanNetworkInventoryView
	return resp, cli.List("v1/l2-networks/vlan", &params, &resp)
}
