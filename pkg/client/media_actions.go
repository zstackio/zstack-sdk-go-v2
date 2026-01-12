// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteMedia deletes Media
func (cli *ZSClient) DeleteMedia(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/media", uuid, string(deleteMode))
}
// QueryMedia queries Media list
func (cli *ZSClient) QueryMedia(params *param.QueryParam) ([]view.MediaInventoryView, error) {
	var resp []view.MediaInventoryView
	return resp, cli.List("v1/media", params, &resp)
}

func (cli *ZSClient) GetMedia(uuid string) (*view.MediaInventoryView, error) {
	var resp view.MediaInventoryView
	if err := cli.Get("v1/media", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
