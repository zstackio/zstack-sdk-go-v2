// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteMedia deletes Media
func (cli *ZSClient) DeleteMedia(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/media/{uuid}", uuid, string(deleteMode))
}
// QueryMedia queries Media list
func (cli *ZSClient) QueryMedia(params *param.QueryParam) ([]view.MediaInventoryView, error) {
	var resp []view.MediaInventoryView
	return resp, cli.List("v1/media", params, &resp)
}
