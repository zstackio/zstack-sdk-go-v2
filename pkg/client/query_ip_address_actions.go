// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIpAddress queries IpAddress list
func (cli *ZSClient) QueryIpAddress(params *param.QueryParam) ([]view.UsedIpInventoryView, error) {
	var resp []view.UsedIpInventoryView
	return resp, cli.List("v1/l3-networks/ip-address", params, &resp)
}
