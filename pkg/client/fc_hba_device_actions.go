// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryFcHbaDevice queries FcHbaDevice list
func (cli *ZSClient) QueryFcHbaDevice(params *param.QueryParam) ([]view.HbaDeviceInventoryView, error) {
	var resp []view.HbaDeviceInventoryView
	return resp, cli.List("v1/storage-devices/hba", params, &resp)
}

// PageFcHbaDevice Pagination
func (cli *ZSClient) PageFcHbaDevice(params *param.QueryParam) ([]view.HbaDeviceInventoryView, int, error) {
	var fcHbaDevices []view.HbaDeviceInventoryView
	total, err := cli.Page("v1/storage-devices/hba", params, &fcHbaDevices)
	return fcHbaDevices, total, err
}
