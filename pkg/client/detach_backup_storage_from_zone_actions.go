// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachBackupStorageFromZone operates on BackupStorageFromZone
func (cli *ZSClient) DetachBackupStorageFromZone(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}", uuid, string(deleteMode))
}
