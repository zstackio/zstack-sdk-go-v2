// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteZone deletes Zone
func (cli *ZSClient) DeleteZone(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zones/{uuid}", uuid, string(deleteMode))
}
