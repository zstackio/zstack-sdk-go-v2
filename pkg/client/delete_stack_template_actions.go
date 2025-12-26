// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteStackTemplate deletes StackTemplate
func (cli *ZSClient) DeleteStackTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cloudformation/template/{uuid}", uuid, string(deleteMode))
}
