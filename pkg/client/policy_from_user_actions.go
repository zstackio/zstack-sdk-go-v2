// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachPolicyFromUser 操作PolicyFromUser
func (cli *ZSClient) DetachPolicyFromUser(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/users/{userUuid}/policies/{policyUuid}", uuid, string(deleteMode))
}

