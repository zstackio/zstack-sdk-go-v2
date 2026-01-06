// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNativeHost queries NativeHost list
func (cli *ZSClient) QueryNativeHost(params *param.QueryParam) ([]view.NativeHostInventoryView, error) {
	var resp []view.NativeHostInventoryView
	return resp, cli.List("v1/container/native/host", params, &resp)
}
