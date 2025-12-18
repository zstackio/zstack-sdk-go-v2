// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveIAM2VirtualIDGroupFromProjects 操作RemoveIAM2VirtualIDGroupFromProjects
func (cli *ZSClient) RemoveIAM2VirtualIDGroupFromProjects(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups", uuid, string(deleteMode))
}

