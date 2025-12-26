// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryL3Network queries L3Network list
func (cli *ZSClient) QueryL3Network(params *param.QueryParam) ([]view.L3NetworkInventoryView, error) {
	var resp []view.L3NetworkInventoryView
	return resp, cli.List("v1/l3-networks", params, &resp)
}
