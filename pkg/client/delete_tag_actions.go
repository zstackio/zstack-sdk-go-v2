// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteTag deletes Tag
func (cli *ZSClient) DeleteTag(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tags/{uuid}", uuid, string(deleteMode))
}
