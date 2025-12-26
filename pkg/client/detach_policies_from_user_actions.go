// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachPoliciesFromUser operates on PoliciesFromUser
func (cli *ZSClient) DetachPoliciesFromUser(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/users/{userUuid}/policies", uuid, string(deleteMode))
}
