// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// ExpungeImageGroup operates on ImageGroup
func (cli *ZSClient) ExpungeImageGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/imagegroups/{uuid}/actions", uuid, string(deleteMode))
}
