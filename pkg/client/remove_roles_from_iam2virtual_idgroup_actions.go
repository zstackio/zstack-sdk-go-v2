// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveRolesFromIAM2VirtualIDGroup 操作RemoveRolesFromIAM2VirtualIDGroup
func (cli *ZSClient) RemoveRolesFromIAM2VirtualIDGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups/{groupUuid}/roles", uuid, string(deleteMode))
}

