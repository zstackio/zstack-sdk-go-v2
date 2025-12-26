// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteImage deletes Image
func (cli *ZSClient) DeleteImage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/images/{uuid}", uuid, string(deleteMode))
}
