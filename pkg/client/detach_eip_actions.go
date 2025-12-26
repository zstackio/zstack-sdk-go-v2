// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachEip operates on Eip
func (cli *ZSClient) DetachEip(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/eips/{uuid}/vm-instances/nics", uuid, string(deleteMode))
}
