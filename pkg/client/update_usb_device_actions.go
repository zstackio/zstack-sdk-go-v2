// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateUsbDevice updates UsbDevice
func (cli *ZSClient) UpdateUsbDevice(uuid string, params param.UpdateUsbDeviceParam) (*view.UpdateUsbDeviceEventView, error) {
	resp := view.UpdateUsbDeviceEventView{}
	if err := cli.Put("v1/usb-device/usb-devices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
