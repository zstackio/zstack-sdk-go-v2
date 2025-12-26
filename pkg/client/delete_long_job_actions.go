// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteLongJob deletes LongJob
func (cli *ZSClient) DeleteLongJob(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/longjobs/{uuid}", uuid, string(deleteMode))
}
