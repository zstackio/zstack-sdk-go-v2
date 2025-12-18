// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachIAM2ProjectFromIAM2Organization 操作IAM2ProjectFromIAM2Organization
func (cli *ZSClient) DetachIAM2ProjectFromIAM2Organization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{projectUuid}/iam2/organizations", uuid, string(deleteMode))
}

