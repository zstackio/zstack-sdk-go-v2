// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachPolicyFromUser operates on PolicyFromUser
func (cli *ZSClient) DetachPolicyFromUser(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/users/{userUuid}/policies/{policyUuid}", uuid, string(deleteMode))
}
