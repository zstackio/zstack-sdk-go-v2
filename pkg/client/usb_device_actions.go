// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryUsbDevice queries UsbDevice list
func (cli *ZSClient) QueryUsbDevice(params *param.QueryParam) ([]view.UsbDeviceInventoryView, error) {
	var resp []view.UsbDeviceInventoryView
	return resp, cli.List("v1/usb-device/usb-devices", params, &resp)
}

// PageUsbDevice Pagination
func (cli *ZSClient) PageUsbDevice(params *param.QueryParam) ([]view.UsbDeviceInventoryView, int, error) {
	var usbDevices []view.UsbDeviceInventoryView
	total, err := cli.Page("v1/usb-device/usb-devices", params, &usbDevices)
	return usbDevices, total, err
}
// UpdateUsbDevice updates UsbDevice
func (cli *ZSClient) UpdateUsbDevice(uuid string, params param.UpdateUsbDeviceParam) (*view.UsbDeviceInventoryView, error) {
	resp := view.UsbDeviceInventoryView{}
	if err := cli.Put("v1/usb-device/usb-devices", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
