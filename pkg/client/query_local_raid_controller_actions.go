// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryLocalRaidController queries LocalRaidController list
func (cli *ZSClient) QueryLocalRaidController(params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, error) {
	var resp []view.RaidPhysicalDriveInventoryView
	return resp, cli.List("v1/storage-devices/local-raid/controllers", params, &resp)
}
