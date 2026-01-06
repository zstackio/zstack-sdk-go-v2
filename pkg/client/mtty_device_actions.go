// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMttyDevice queries MttyDevice list
func (cli *ZSClient) QueryMttyDevice(params *param.QueryParam) ([]view.MttyDeviceInventoryView, error) {
	var resp []view.MttyDeviceInventoryView
	return resp, cli.List("v1/mtty-devices", params, &resp)
}
