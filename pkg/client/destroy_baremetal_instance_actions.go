// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DestroyBaremetalInstance destroys BaremetalInstance
func (cli *ZSClient) DestroyBaremetalInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/instances/{uuid}", uuid, string(deleteMode))
}
