// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachUsbDeviceFromVm 操作UsbDeviceFromVm
func (cli *ZSClient) DetachUsbDeviceFromVm(params param.DetachUsbDeviceFromVmParam) (*view.DetachUsbDeviceFromVmEventView, error) {
	resp := view.DetachUsbDeviceFromVmEventView{}
	if err := cli.Post("v1/usb-device/usb-devices/{usbDeviceUuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

