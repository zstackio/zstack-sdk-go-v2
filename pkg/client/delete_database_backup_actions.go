// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteDatabaseBackup deletes DatabaseBackup
func (cli *ZSClient) DeleteDatabaseBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/database-backups/{uuid}", uuid, string(deleteMode))
}
