// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteDatabaseBackup deletes DatabaseBackup
func (cli *ZSClient) DeleteDatabaseBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/database-backups/{uuid}", uuid, string(deleteMode))
}
