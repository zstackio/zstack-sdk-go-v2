// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryFcHbaDevice queries FcHbaDevice list
func (cli *ZSClient) QueryFcHbaDevice(params param.QueryParam) ([]view.HbaDeviceInventoryView, error) {
	var resp []view.HbaDeviceInventoryView
	return resp, cli.List("v1/storage-devices/hba", &params, &resp)
}
