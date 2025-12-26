// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteExternalBackup deletes ExternalBackup
func (cli *ZSClient) DeleteExternalBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/externalbackup/{uuid}", uuid, string(deleteMode))
}
