// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveMonFromCephPrimaryStorage removes MonFromCephPrimaryStorage
func (cli *ZSClient) RemoveMonFromCephPrimaryStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/ceph/{uuid}/mons", uuid, string(deleteMode))
}
