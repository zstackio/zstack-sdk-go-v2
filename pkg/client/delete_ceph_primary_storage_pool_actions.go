// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteCephPrimaryStoragePool deletes CephPrimaryStoragePool
func (cli *ZSClient) DeleteCephPrimaryStoragePool(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/ceph/pools/{uuid}", uuid, string(deleteMode))
}
