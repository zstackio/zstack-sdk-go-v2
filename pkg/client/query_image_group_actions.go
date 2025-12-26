// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryImageGroup queries ImageGroup list
func (cli *ZSClient) QueryImageGroup(params *param.QueryParam) ([]view.ImageGroupInventoryView, error) {
	var resp []view.ImageGroupInventoryView
	return resp, cli.List("v1/imagegroups", params, &resp)
}
