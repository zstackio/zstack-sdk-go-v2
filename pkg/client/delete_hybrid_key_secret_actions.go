// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteHybridKeySecret deletes HybridKeySecret
func (cli *ZSClient) DeleteHybridKeySecret(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/hybrid/key/{uuid}", uuid, string(deleteMode))
}
