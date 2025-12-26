// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEcsVpcInLocal deletes EcsVpcInLocal
func (cli *ZSClient) DeleteEcsVpcInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vpc/{uuid}", uuid, string(deleteMode))
}
