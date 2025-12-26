// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPhysicalDriveSelfTestHistory queries PhysicalDriveSelfTestHistory list
func (cli *ZSClient) QueryPhysicalDriveSelfTestHistory(params *param.QueryParam) ([]view.PhysicalDriveSmartSelfTestHistoryInventoryView, error) {
	var resp []view.PhysicalDriveSmartSelfTestHistoryInventoryView
	return resp, cli.List("v1/storage-devices/local-raid/physical-drives/self-test", params, &resp)
}
