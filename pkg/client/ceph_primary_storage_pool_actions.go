// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteCephPrimaryStoragePool 删除CephPrimaryStoragePool
func (cli *ZSClient) DeleteCephPrimaryStoragePool(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/ceph/pools/{uuid}", uuid, string(deleteMode))
}

