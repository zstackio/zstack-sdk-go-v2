// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveAttributesFromIAM2Organization 操作RemoveAttributesFromIAM2Organization
func (cli *ZSClient) RemoveAttributesFromIAM2Organization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/organizations/{uuid}/attributes", uuid, string(deleteMode))
}

