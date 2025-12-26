// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryZBox queries ZBox list
func (cli *ZSClient) QueryZBox(params *param.QueryParam) ([]view.ZBoxInventoryView, error) {
	var resp []view.ZBoxInventoryView
	return resp, cli.List("v1/zbox", params, &resp)
}
