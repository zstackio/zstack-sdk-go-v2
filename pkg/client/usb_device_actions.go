// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryUsbDevice queries UsbDevice list
func (cli *ZSClient) QueryUsbDevice(params *param.QueryParam) ([]view.UsbDeviceInventoryView, error) {
	var resp []view.UsbDeviceInventoryView
	return resp, cli.List("v1/usb-device/usb-devices", params, &resp)
}
// UpdateUsbDevice updates UsbDevice
func (cli *ZSClient) UpdateUsbDevice(uuid string, params param.UpdateUsbDeviceParam) (*view.UsbDeviceInventoryView, error) {
	var resp view.UpdateUsbDeviceEventView
	if err := cli.Put("v1/usb-device/usb-devices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
