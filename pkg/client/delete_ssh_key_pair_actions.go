// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSshKeyPair deletes SshKeyPair
func (cli *ZSClient) DeleteSshKeyPair(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ssh-key-pair/{uuid}", uuid, string(deleteMode))
}
