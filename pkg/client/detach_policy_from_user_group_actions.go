// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachPolicyFromUserGroup operates on PolicyFromUserGroup
func (cli *ZSClient) DetachPolicyFromUserGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/groups/{groupUuid}/policies/{policyUuid}", uuid, string(deleteMode))
}
