// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSecurityMachine deletes SecurityMachine
func (cli *ZSClient) DeleteSecurityMachine(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-machines/{uuid}", uuid, string(deleteMode))
}
