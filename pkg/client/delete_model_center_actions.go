// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteModelCenter deletes ModelCenter
func (cli *ZSClient) DeleteModelCenter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-centers/{uuid}", uuid, string(deleteMode))
}
