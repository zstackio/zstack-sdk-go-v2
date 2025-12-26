// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteIAM2VirtualID deletes IAM2VirtualID
func (cli *ZSClient) DeleteIAM2VirtualID(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/virtual-ids/{uuid}", uuid, string(deleteMode))
}
