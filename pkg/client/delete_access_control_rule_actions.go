// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAccessControlRule deletes AccessControlRule
func (cli *ZSClient) DeleteAccessControlRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/login-control/access-control/rules/{uuid}", uuid, string(deleteMode))
}
