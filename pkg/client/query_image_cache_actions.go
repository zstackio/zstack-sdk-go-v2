// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryImageCache queries ImageCache list
func (cli *ZSClient) QueryImageCache(params *param.QueryParam) ([]view.ImageCacheInventoryView, error) {
	var resp []view.ImageCacheInventoryView
	return resp, cli.List("v1/primary-storage/imagecache", params, &resp)
}
