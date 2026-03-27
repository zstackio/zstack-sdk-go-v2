// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ExpungeImageGroup operates on ImageGroup
func (cli *ZSClient) ExpungeImageGroup(ctx context.Context, uuid string) error {
	params := map[string]interface{}{
		"expungeImageGroup": map[string]interface{}{},
	}
	return cli.Put(ctx, "v1/imagegroups", uuid, params, nil)
}
// QueryImageGroup queries ImageGroup list
func (cli *ZSClient) QueryImageGroup(ctx context.Context, params *param.QueryParam) ([]view.ImageGroupInventoryView, error) {
	var resp []view.ImageGroupInventoryView
	return resp, cli.List(ctx, "v1/imagegroups", params, &resp)
}

func (cli *ZSClient) GetImageGroup(ctx context.Context, uuid string) (*view.ImageGroupInventoryView, error) {
	var resp view.ImageGroupInventoryView
	if err := cli.Get(ctx, "v1/imagegroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImageGroup Pagination
func (cli *ZSClient) PageImageGroup(ctx context.Context, params *param.QueryParam) ([]view.ImageGroupInventoryView, int, error) {
	var imageGroups []view.ImageGroupInventoryView
	total, err := cli.Page(ctx, "v1/imagegroups", params, &imageGroups)
	return imageGroups, total, err
}
