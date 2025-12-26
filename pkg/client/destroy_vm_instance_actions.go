// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DestroyVmInstance destroys VmInstance
func (cli *ZSClient) DestroyVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}", uuid, string(deleteMode))
}
