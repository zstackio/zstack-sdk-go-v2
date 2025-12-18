// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVirtualRouterOffering queries VirtualRouterOffering list
func (cli *ZSClient) QueryVirtualRouterOffering(params param.QueryParam) ([]view.VirtualRouterOfferingInventoryView, error) {
	var resp []view.VirtualRouterOfferingInventoryView
	return resp, cli.List("v1/instance-offerings/virtual-routers", &params, &resp)
}
