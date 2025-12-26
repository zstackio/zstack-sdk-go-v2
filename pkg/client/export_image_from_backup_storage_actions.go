// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ExportImageFromBackupStorage operates on ExportImageFromBackupStorage
func (cli *ZSClient) ExportImageFromBackupStorage(uuid string, params param.ExportImageFromBackupStorageParam) (*view.ExportImageFromBackupStorageEventView, error) {
	resp := view.ExportImageFromBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
