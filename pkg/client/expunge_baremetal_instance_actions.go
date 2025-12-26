// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// ExpungeBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) ExpungeBaremetalInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/instances/{uuid}/actions", uuid, string(deleteMode))
}
