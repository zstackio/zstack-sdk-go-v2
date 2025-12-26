// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveAttributesFromIAM2VirtualID removes AttributesFromIAM2VirtualID
func (cli *ZSClient) RemoveAttributesFromIAM2VirtualID(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/virtual-ids/{uuid}/attributes", uuid, string(deleteMode))
}
