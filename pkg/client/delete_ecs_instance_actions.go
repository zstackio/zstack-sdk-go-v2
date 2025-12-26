// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEcsInstance deletes EcsInstance
func (cli *ZSClient) DeleteEcsInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/ecs/{uuid}/remote", uuid, string(deleteMode))
}
