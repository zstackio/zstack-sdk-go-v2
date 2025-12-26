// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveAttributesFromIAM2Project removes AttributesFromIAM2Project
func (cli *ZSClient) RemoveAttributesFromIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{uuid}/attributes", uuid, string(deleteMode))
}
