// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ExpungeImageGroup operates on ImageGroup
func (cli *ZSClient) ExpungeImageGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/imagegroups/{uuid}/actions", uuid, string(deleteMode))
}
// QueryImageGroup queries ImageGroup list
func (cli *ZSClient) QueryImageGroup(params *param.QueryParam) ([]view.ImageGroupInventoryView, error) {
	var resp []view.ImageGroupInventoryView
	return resp, cli.List("v1/imagegroups", params, &resp)
}
