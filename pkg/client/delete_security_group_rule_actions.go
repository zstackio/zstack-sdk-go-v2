// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSecurityGroupRule deletes SecurityGroupRule
func (cli *ZSClient) DeleteSecurityGroupRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups/rules", uuid, string(deleteMode))
}
