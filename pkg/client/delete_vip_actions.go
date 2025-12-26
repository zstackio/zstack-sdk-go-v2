// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVip deletes Vip
func (cli *ZSClient) DeleteVip(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vips/{uuid}", uuid, string(deleteMode))
}
