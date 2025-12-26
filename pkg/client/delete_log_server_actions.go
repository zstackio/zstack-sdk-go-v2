// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteLogServer deletes LogServer
func (cli *ZSClient) DeleteLogServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/log/servers", uuid, string(deleteMode))
}
