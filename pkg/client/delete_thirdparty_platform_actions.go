// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteThirdpartyPlatform deletes ThirdpartyPlatform
func (cli *ZSClient) DeleteThirdpartyPlatform(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/third-party/platforms/{uuid}", uuid, string(deleteMode))
}
