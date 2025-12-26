// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEcsImageLocal deletes EcsImageLocal
func (cli *ZSClient) DeleteEcsImageLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/image/{uuid}", uuid, string(deleteMode))
}
