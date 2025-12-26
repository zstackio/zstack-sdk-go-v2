// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryImagePackage queries ImagePackage list
func (cli *ZSClient) QueryImagePackage(params *param.QueryParam) ([]view.ImagePackageInventoryView, error) {
	var resp []view.ImagePackageInventoryView
	return resp, cli.List("v1/image-packages", params, &resp)
}
