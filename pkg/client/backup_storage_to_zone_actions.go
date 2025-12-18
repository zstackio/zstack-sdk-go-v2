// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachBackupStorageToZone 操作BackupStorageToZone
func (cli *ZSClient) AttachBackupStorageToZone(params param.AttachBackupStorageToZoneParam) (*view.AttachBackupStorageToZoneEventView, error) {
	resp := view.AttachBackupStorageToZoneEventView{}
	if err := cli.Post("v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

