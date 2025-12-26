// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteHybridEipRemote deletes HybridEipRemote
func (cli *ZSClient) DeleteHybridEipRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/eip/{uuid}/remote", uuid, string(deleteMode))
}
