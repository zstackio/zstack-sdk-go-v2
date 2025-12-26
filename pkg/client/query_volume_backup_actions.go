// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVolumeBackup queries VolumeBackup list
func (cli *ZSClient) QueryVolumeBackup(params *param.QueryParam) ([]view.VolumeBackupInventoryView, error) {
	var resp []view.VolumeBackupInventoryView
	return resp, cli.List("v1/volume-backups", params, &resp)
}
