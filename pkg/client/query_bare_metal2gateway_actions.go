// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBareMetal2Gateway queries BareMetal2Gateway list
func (cli *ZSClient) QueryBareMetal2Gateway(params *param.QueryParam) ([]view.BareMetal2GatewayInventoryView, error) {
	var resp []view.BareMetal2GatewayInventoryView
	return resp, cli.List("v1/baremetal2/gateways", params, &resp)
}
