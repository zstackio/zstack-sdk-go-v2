// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryUsbDevice queries UsbDevice list
func (cli *ZSClient) QueryUsbDevice(ctx context.Context, params *param.QueryParam) ([]view.UsbDeviceInventoryView, error) {
	var resp []view.UsbDeviceInventoryView
	return resp, cli.List(ctx, "v1/usb-device/usb-devices", params, &resp)
}

func (cli *ZSClient) GetUsbDevice(ctx context.Context, uuid string) (*view.UsbDeviceInventoryView, error) {
	var resp view.UsbDeviceInventoryView
	if err := cli.Get(ctx, "v1/usb-device/usb-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageUsbDevice Pagination
func (cli *ZSClient) PageUsbDevice(ctx context.Context, params *param.QueryParam) ([]view.UsbDeviceInventoryView, int, error) {
	var usbDevices []view.UsbDeviceInventoryView
	total, err := cli.Page(ctx, "v1/usb-device/usb-devices", params, &usbDevices)
	return usbDevices, total, err
}
// UpdateUsbDevice updates UsbDevice
func (cli *ZSClient) UpdateUsbDevice(ctx context.Context, uuid string, params param.UpdateUsbDeviceParam) (*view.UsbDeviceInventoryView, error) {
	resp := view.UsbDeviceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/usb-device/usb-devices", uuid, "", map[string]interface{}{
		"updateUsbDevice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
