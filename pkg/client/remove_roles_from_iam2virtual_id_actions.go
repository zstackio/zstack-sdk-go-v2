// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveRolesFromIAM2VirtualID 操作RemoveRolesFromIAM2VirtualID
func (cli *ZSClient) RemoveRolesFromIAM2VirtualID(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/virtual-ids/{virtualIDUuid}/roles", uuid, string(deleteMode))
}

