// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ExpungeImageGroup operates on ImageGroup
func (cli *ZSClient) ExpungeImageGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/imagegroups", fmt.Sprintf(\"%s/actions\", uuid), string(deleteMode))
}
// QueryImageGroup queries ImageGroup list
func (cli *ZSClient) QueryImageGroup(params *param.QueryParam) ([]view.ImageGroupInventoryView, error) {
	var resp []view.ImageGroupInventoryView
	return resp, cli.List("v1/imagegroups", params, &resp)
}

func (cli *ZSClient) GetImageGroup(uuid string) (*view.ImageGroupInventoryView, error) {
	var resp view.ImageGroupInventoryView
	if err := cli.Get("v1/imagegroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
