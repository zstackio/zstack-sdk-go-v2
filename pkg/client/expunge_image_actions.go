// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// ExpungeImage operates on Image
func (cli *ZSClient) ExpungeImage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/images/{imageUuid}/actions", uuid, string(deleteMode))
}
