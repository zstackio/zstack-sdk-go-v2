// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveAttributesFromIAM2VirtualIDGroup 操作RemoveAttributesFromIAM2VirtualIDGroup
func (cli *ZSClient) RemoveAttributesFromIAM2VirtualIDGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups/{uuid}/attributes", uuid, string(deleteMode))
}

