// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteImagePackage deletes ImagePackage
func (cli *ZSClient) DeleteImagePackage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/image-packages/{uuid}", uuid, string(deleteMode))
}
