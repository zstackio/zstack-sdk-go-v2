// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBaremetalPxeServer queries BaremetalPxeServer list
func (cli *ZSClient) QueryBaremetalPxeServer(params *param.QueryParam) ([]view.BaremetalPxeServerInventoryView, error) {
	var resp []view.BaremetalPxeServerInventoryView
	return resp, cli.List("v1/baremetal/pxeservers", params, &resp)
}
