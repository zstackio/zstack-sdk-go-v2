// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryImageGroupRef queries ImageGroupRef list
func (cli *ZSClient) QueryImageGroupRef(params *param.QueryParam) ([]view.ImageGroupRefInventoryView, error) {
	var resp []view.ImageGroupRefInventoryView
	return resp, cli.List("v1/imagegrouprefs", params, &resp)
}
