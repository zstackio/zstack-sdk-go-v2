// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveRolesFromIAM2VirtualID removes RolesFromIAM2VirtualID
func (cli *ZSClient) RemoveRolesFromIAM2VirtualID(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/virtual-ids/{virtualIDUuid}/roles", uuid, string(deleteMode))
}
