// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEcsImageRemote deletes EcsImageRemote
func (cli *ZSClient) DeleteEcsImageRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/image/remote/{uuid}", uuid, string(deleteMode))
}
