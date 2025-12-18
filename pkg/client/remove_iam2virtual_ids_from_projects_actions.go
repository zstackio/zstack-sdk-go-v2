// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveIAM2VirtualIDsFromProjects 操作RemoveIAM2VirtualIDsFromProjects
func (cli *ZSClient) RemoveIAM2VirtualIDsFromProjects(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/virtual-ids", uuid, string(deleteMode))
}

