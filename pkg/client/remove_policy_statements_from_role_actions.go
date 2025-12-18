// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemovePolicyStatementsFromRole removes PolicyStatementsFromRole
func (cli *ZSClient) RemovePolicyStatementsFromRole(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/roles/{uuid}/policy-statements", uuid, string(deleteMode))
}
