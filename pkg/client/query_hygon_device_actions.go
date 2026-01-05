// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryHygonDevice queries HygonDevice list
func (cli *ZSClient) QueryHygonDevice(params *param.QueryParam) ([]view.HygonCcpDeviceInventoryView, error) {
	var resp []view.HygonCcpDeviceInventoryView
	return resp, cli.List("v1/hygon-devices", params, &resp)
}
