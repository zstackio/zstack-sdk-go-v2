// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteExportedImageFromBackupStorage 删除ExportedImageFromBackupStorage
func (cli *ZSClient) DeleteExportedImageFromBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/backup-storage/{backupStorageUuid}/exported-images/{imageUuid}", uuid, string(deleteMode))
}

