// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteMedia deletes Media
func (cli *ZSClient) DeleteMedia(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/media/{uuid}", uuid, string(deleteMode))
}
