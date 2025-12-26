// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachSshKeyPairFromVmInstance operates on SshKeyPairFromVmInstance
func (cli *ZSClient) DetachSshKeyPairFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ssh-key-pair/{sshKeyPairUuid}/vm-instance/{vmInstanceUuid}", uuid, string(deleteMode))
}
