// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteUser 删除User
func (cli *ZSClient) DeleteUser(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/users/{uuid}", uuid, string(deleteMode))
}

