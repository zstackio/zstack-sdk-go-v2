// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteResourceStack deletes ResourceStack
func (cli *ZSClient) DeleteResourceStack(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cloudformation/stack/{uuid}", uuid, string(deleteMode))
}
