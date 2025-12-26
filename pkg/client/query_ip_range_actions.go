// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIpRange queries IpRange list
func (cli *ZSClient) QueryIpRange(params *param.QueryParam) ([]view.IpRangeInventoryView, error) {
	var resp []view.IpRangeInventoryView
	return resp, cli.List("v1/l3-networks/ip-ranges", params, &resp)
}
