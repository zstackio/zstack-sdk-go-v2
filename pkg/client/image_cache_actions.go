// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryImageCache queries ImageCache list
func (cli *ZSClient) QueryImageCache(params *param.QueryParam) ([]view.ImageCacheInventoryView, error) {
	var resp []view.ImageCacheInventoryView
	return resp, cli.List("v1/primary-storage/imagecache", params, &resp)
}

func (cli *ZSClient) GetImageCache(uuid string) (*view.ImageCacheInventoryView, error) {
	var resp view.ImageCacheInventoryView
	if err := cli.Get("v1/primary-storage/imagecache", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
