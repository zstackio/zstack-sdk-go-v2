// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveAttributesFromIAM2Project 操作RemoveAttributesFromIAM2Project
func (cli *ZSClient) RemoveAttributesFromIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{uuid}/attributes", uuid, string(deleteMode))
}

