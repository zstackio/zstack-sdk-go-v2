// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteSecretResourcePool deletes SecretResourcePool
func (cli *ZSClient) DeleteSecretResourcePool(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/secret-resource-pool/{uuid}", uuid, string(deleteMode))
}
