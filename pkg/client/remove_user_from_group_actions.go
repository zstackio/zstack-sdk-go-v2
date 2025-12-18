// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveUserFromGroup 操作RemoveUserFromGroup
func (cli *ZSClient) RemoveUserFromGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/groups/{groupUuid}/users/{userUuid}", uuid, string(deleteMode))
}

