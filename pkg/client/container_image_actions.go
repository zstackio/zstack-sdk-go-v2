// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryContainerImage queries ContainerImage list
func (cli *ZSClient) QueryContainerImage(params *param.QueryParam) ([]view.ContainerImageInventoryView, error) {
	var resp []view.ContainerImageInventoryView
	return resp, cli.List("v1/container/images", params, &resp)
}
