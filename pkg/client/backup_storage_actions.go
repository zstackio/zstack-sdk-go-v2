// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteBackupStorage 删除BackupStorage
func (cli *ZSClient) DeleteBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/backup-storage/{uuid}", uuid, string(deleteMode))
}

