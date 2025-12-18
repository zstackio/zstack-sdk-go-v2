// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryImageCache 查询ImageCache列表
func (cli *ZSClient) QueryImageCache(params param.QueryParam) ([]view.QueryImageCacheView, error) {
	var resp []view.QueryImageCacheView
	return resp, cli.List("v1/primary-storage/imagecache", &params, &resp)
}

