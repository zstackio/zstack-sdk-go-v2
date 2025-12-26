// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveAttributesFromIAM2Organization removes AttributesFromIAM2Organization
func (cli *ZSClient) RemoveAttributesFromIAM2Organization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/organizations/{uuid}/attributes", uuid, string(deleteMode))
}
