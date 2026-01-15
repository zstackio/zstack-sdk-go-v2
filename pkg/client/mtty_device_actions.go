// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMttyDevice queries MttyDevice list
func (cli *ZSClient) QueryMttyDevice(params *param.QueryParam) ([]view.MttyDeviceInventoryView, error) {
	var resp []view.MttyDeviceInventoryView
	return resp, cli.List("v1/mtty-devices", params, &resp)
}

func (cli *ZSClient) GetMttyDevice(uuid string) (*view.MttyDeviceInventoryView, error) {
	var resp view.MttyDeviceInventoryView
	if err := cli.Get("v1/mtty-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMttyDevice Pagination
func (cli *ZSClient) PageMttyDevice(params *param.QueryParam) ([]view.MttyDeviceInventoryView, int, error) {
	var mttyDevices []view.MttyDeviceInventoryView
	total, err := cli.Page("v1/mtty-devices", params, &mttyDevices)
	return mttyDevices, total, err
}
