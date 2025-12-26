// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveMonFromCephBackupStorage removes MonFromCephBackupStorage
func (cli *ZSClient) RemoveMonFromCephBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/backup-storage/ceph/{uuid}/mons", uuid, string(deleteMode))
}
