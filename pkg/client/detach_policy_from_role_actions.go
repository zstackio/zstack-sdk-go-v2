// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachPolicyFromRole operates on PolicyFromRole
func (cli *ZSClient) DetachPolicyFromRole(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/policies/{policyUuid}/roles/{roleUuid}", uuid, string(deleteMode))
}
