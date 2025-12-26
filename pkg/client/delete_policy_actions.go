// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePolicy deletes Policy
func (cli *ZSClient) DeletePolicy(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/policies/{uuid}", uuid, string(deleteMode))
}
