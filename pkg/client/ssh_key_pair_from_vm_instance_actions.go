// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachSshKeyPairFromVmInstance 操作SshKeyPairFromVmInstance
func (cli *ZSClient) DetachSshKeyPairFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ssh-key-pair/{sshKeyPairUuid}/vm-instance/{vmInstanceUuid}", uuid, string(deleteMode))
}

