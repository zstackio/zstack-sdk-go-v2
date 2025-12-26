// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryZone queries Zone list
func (cli *ZSClient) QueryZone(params *param.QueryParam) ([]view.ZoneInventoryView, error) {
	var resp []view.ZoneInventoryView
	return resp, cli.List("v1/zones", params, &resp)
}
