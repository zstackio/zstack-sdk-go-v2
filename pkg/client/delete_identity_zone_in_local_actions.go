// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteIdentityZoneInLocal deletes IdentityZoneInLocal
func (cli *ZSClient) DeleteIdentityZoneInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/identity-zone/{uuid}", uuid, string(deleteMode))
}
