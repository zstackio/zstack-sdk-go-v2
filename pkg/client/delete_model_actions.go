// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteModel deletes Model
func (cli *ZSClient) DeleteModel(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/models/{uuid}", uuid, string(deleteMode))
}
