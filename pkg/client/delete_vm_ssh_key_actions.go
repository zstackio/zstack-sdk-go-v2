// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmSshKey deletes VmSshKey
func (cli *ZSClient) DeleteVmSshKey(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/ssh-keys", uuid, string(deleteMode))
}
