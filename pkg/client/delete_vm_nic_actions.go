// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmNic deletes VmNic
func (cli *ZSClient) DeleteVmNic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nics/{uuid}", uuid, string(deleteMode))
}
