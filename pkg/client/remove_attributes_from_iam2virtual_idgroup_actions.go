// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveAttributesFromIAM2VirtualIDGroup removes AttributesFromIAM2VirtualIDGroup
func (cli *ZSClient) RemoveAttributesFromIAM2VirtualIDGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups/{uuid}/attributes", uuid, string(deleteMode))
}
