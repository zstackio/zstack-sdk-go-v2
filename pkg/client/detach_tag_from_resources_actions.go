// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachTagFromResources operates on TagFromResources
func (cli *ZSClient) DetachTagFromResources(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tags/{tagUuid}/resources", uuid, string(deleteMode))
}
