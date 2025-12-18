// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVmBackup 删除VmBackup
func (cli *ZSClient) DeleteVmBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-backups/{groupUuid}", uuid, string(deleteMode))
}

