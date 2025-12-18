// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteIAM2Project 删除IAM2Project
func (cli *ZSClient) DeleteIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{uuid}", uuid, string(deleteMode))
}

