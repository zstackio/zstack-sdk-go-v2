// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteConnectionAccessPointLocal deletes ConnectionAccessPointLocal
func (cli *ZSClient) DeleteConnectionAccessPointLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/access-point/{uuid}", uuid, string(deleteMode))
}
