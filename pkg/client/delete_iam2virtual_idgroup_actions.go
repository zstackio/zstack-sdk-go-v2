// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteIAM2VirtualIDGroup deletes IAM2VirtualIDGroup
func (cli *ZSClient) DeleteIAM2VirtualIDGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups/{uuid}", uuid, string(deleteMode))
}
