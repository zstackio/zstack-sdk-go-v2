// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSecretResourcePool deletes SecretResourcePool
func (cli *ZSClient) DeleteSecretResourcePool(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/secret-resource-pool/{uuid}", uuid, string(deleteMode))
}
