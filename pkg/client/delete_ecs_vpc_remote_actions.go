// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEcsVpcRemote deletes EcsVpcRemote
func (cli *ZSClient) DeleteEcsVpcRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vpc/remote/{uuid}", uuid, string(deleteMode))
}
