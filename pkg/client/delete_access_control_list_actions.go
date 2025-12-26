// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAccessControlList deletes AccessControlList
func (cli *ZSClient) DeleteAccessControlList(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/access-control-lists/{uuid}", uuid, string(deleteMode))
}
