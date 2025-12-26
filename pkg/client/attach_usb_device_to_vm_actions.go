// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachUsbDeviceToVm operates on UsbDeviceToVm
func (cli *ZSClient) AttachUsbDeviceToVm(params param.AttachUsbDeviceToVmParam) (*view.AttachUsbDeviceToVmEventView, error) {
	resp := view.AttachUsbDeviceToVmEventView{}
	if err := cli.Post("v1/usb-device/usb-devices/{usbDeviceUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
