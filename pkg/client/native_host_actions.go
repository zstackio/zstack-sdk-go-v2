// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNativeHost queries NativeHost list
func (cli *ZSClient) QueryNativeHost(params *param.QueryParam) ([]view.NativeHostInventoryView, error) {
	var resp []view.NativeHostInventoryView
	return resp, cli.List("v1/container/native/host", params, &resp)
}

// PageNativeHost Pagination
func (cli *ZSClient) PageNativeHost(params *param.QueryParam) ([]view.NativeHostInventoryView, int, error) {
	var nativeHosts []view.NativeHostInventoryView
	total, err := cli.Page("v1/container/native/host", params, &nativeHosts)
	return nativeHosts, total, err
}
