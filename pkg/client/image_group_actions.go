// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ExpungeImageGroup operates on ImageGroup
func (cli *ZSClient) ExpungeImageGroup(uuid string) error {
	params := map[string]interface{}{
		"expungeImageGroup": map[string]interface{}{},
	}
	return cli.Put("v1/imagegroups", uuid, params, nil)
}
// QueryImageGroup queries ImageGroup list
func (cli *ZSClient) QueryImageGroup(params *param.QueryParam) ([]view.ImageGroupInventoryView, error) {
	var resp []view.ImageGroupInventoryView
	return resp, cli.List("v1/imagegroups", params, &resp)
}

// PageImageGroup Pagination
func (cli *ZSClient) PageImageGroup(params *param.QueryParam) ([]view.ImageGroupInventoryView, int, error) {
	var imageGroups []view.ImageGroupInventoryView
	total, err := cli.Page("v1/imagegroups", params, &imageGroups)
	return imageGroups, total, err
}
