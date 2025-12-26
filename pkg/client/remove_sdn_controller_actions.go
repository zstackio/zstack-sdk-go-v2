// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveSdnController removes SdnController
func (cli *ZSClient) RemoveSdnController(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controllers/{uuid}", uuid, string(deleteMode))
}
