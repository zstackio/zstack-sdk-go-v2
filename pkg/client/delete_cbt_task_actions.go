// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteCbtTask deletes CbtTask
func (cli *ZSClient) DeleteCbtTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cbt-task/{uuid}", uuid, string(deleteMode))
}
