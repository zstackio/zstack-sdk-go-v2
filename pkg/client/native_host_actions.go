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

func (cli *ZSClient) GetNativeHost(uuid string) (*view.NativeHostInventoryView, error) {
	var resp view.NativeHostInventoryView
	if err := cli.Get("v1/container/native/host", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
