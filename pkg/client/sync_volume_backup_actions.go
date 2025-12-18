// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncVolumeBackup 操作SyncVolumeBackup
func (cli *ZSClient) SyncVolumeBackup(uuid string, params param.SyncVolumeBackupParam) (*view.SyncVolumeBackupEventView, error) {
	resp := view.SyncVolumeBackupEventView{}
	if err := cli.Put("v1/volume-backups/imageStore/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

