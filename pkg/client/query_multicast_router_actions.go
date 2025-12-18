// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMulticastRouter queries MulticastRouter list
func (cli *ZSClient) QueryMulticastRouter(params param.QueryParam) ([]view.MulticastRouterInventoryView, error) {
	var resp []view.MulticastRouterInventoryView
	return resp, cli.List("v1/multicast/virtual-routers", &params, &resp)
}
