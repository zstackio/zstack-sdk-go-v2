// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryFcHbaDevice queries FcHbaDevice list
func (cli *ZSClient) QueryFcHbaDevice(params *param.QueryParam) ([]view.HbaDeviceInventoryView, error) {
	var resp []view.HbaDeviceInventoryView
	return resp, cli.List("v1/storage-devices/hba", params, &resp)
}
