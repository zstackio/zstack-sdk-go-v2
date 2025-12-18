// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryUsbDevice 查询UsbDevice列表
func (cli *ZSClient) QueryUsbDevice(params param.QueryParam) ([]view.QueryUsbDeviceView, error) {
	var resp []view.QueryUsbDeviceView
	return resp, cli.List("v1/usb-device/usb-devices", &params, &resp)
}

