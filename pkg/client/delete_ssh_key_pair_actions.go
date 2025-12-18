// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteSshKeyPair deletes SshKeyPair
func (cli *ZSClient) DeleteSshKeyPair(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ssh-key-pair/{uuid}", uuid, string(deleteMode))
}
