// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVCenter deletes VCenter
func (cli *ZSClient) DeleteVCenter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vcenters/{uuid}", uuid, string(deleteMode))
}
