// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMdevDeviceSpec queries MdevDeviceSpec list
func (cli *ZSClient) QueryMdevDeviceSpec(params *param.QueryParam) ([]view.MdevDeviceSpecInventoryView, error) {
	var resp []view.MdevDeviceSpecInventoryView
	return resp, cli.List("v1/mdev-device-specs", params, &resp)
}
