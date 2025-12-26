// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteModels deletes Models
func (cli *ZSClient) DeleteModels(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/models", uuid, string(deleteMode))
}
