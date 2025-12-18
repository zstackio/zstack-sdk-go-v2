// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryL2VxlanNetworkPool queries L2VxlanNetworkPool list
func (cli *ZSClient) QueryL2VxlanNetworkPool(params param.QueryParam) ([]view.L2VxlanNetworkPoolInventoryView, error) {
	var resp []view.L2VxlanNetworkPoolInventoryView
	return resp, cli.List("v1/l2-networks/vxlan-pool", &params, &resp)
}
