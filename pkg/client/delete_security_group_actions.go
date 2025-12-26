// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSecurityGroup deletes SecurityGroup
func (cli *ZSClient) DeleteSecurityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups/{uuid}", uuid, string(deleteMode))
}
