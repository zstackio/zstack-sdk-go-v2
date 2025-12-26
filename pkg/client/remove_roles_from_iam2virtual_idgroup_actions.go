// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveRolesFromIAM2VirtualIDGroup removes RolesFromIAM2VirtualIDGroup
func (cli *ZSClient) RemoveRolesFromIAM2VirtualIDGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups/{groupUuid}/roles", uuid, string(deleteMode))
}
