// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// ExpungeVmInstance operates on VmInstance
func (cli *ZSClient) ExpungeVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/actions", uuid, string(deleteMode))
}
