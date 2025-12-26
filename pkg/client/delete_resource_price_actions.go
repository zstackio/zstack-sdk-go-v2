// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteResourcePrice deletes ResourcePrice
func (cli *ZSClient) DeleteResourcePrice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/prices/{uuid}", uuid, string(deleteMode))
}
