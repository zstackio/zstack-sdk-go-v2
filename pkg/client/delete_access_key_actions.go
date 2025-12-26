// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAccessKey deletes AccessKey
func (cli *ZSClient) DeleteAccessKey(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accesskeys/{uuid}", uuid, string(deleteMode))
}
