// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEcsInstanceLocal deletes EcsInstanceLocal
func (cli *ZSClient) DeleteEcsInstanceLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/ecs/{uuid}", uuid, string(deleteMode))
}
