// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachBackupStorageToZone operates on BackupStorageToZone
func (cli *ZSClient) AttachBackupStorageToZone(params param.AttachBackupStorageToZoneParam) (*view.AttachBackupStorageToZoneEventView, error) {
	resp := view.AttachBackupStorageToZoneEventView{}
	if err := cli.Post("v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
