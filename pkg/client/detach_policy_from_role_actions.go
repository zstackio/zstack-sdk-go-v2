// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachPolicyFromRole operates on PolicyFromRole
func (cli *ZSClient) DetachPolicyFromRole(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/policies/{policyUuid}/roles/{roleUuid}", uuid, string(deleteMode))
}
