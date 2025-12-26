// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveUserFromGroup removes UserFromGroup
func (cli *ZSClient) RemoveUserFromGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/groups/{groupUuid}/users/{userUuid}", uuid, string(deleteMode))
}
