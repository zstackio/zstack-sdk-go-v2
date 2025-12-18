// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachBackupStorageFromZone 操作BackupStorageFromZone
func (cli *ZSClient) DetachBackupStorageFromZone(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}", uuid, string(deleteMode))
}

