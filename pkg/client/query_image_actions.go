// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryImage queries Image list
func (cli *ZSClient) QueryImage(params *param.QueryParam) ([]view.ImageInventoryView, error) {
	var resp []view.ImageInventoryView
	return resp, cli.List("v1/images", params, &resp)
}
