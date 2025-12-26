// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryUsbDevice queries UsbDevice list
func (cli *ZSClient) QueryUsbDevice(params *param.QueryParam) ([]view.UsbDeviceInventoryView, error) {
	var resp []view.UsbDeviceInventoryView
	return resp, cli.List("v1/usb-device/usb-devices", params, &resp)
}
