// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteBackupFileInPublic 删除BackupFileInPublic
func (cli *ZSClient) DeleteBackupFileInPublic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/backup-mysql", uuid, string(deleteMode))
}

