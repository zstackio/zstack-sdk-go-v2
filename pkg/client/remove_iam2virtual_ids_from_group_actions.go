// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveIAM2VirtualIDsFromGroup removes IAM2VirtualIDsFromGroup
func (cli *ZSClient) RemoveIAM2VirtualIDsFromGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups/{groupUuid}/virtual-ids", uuid, string(deleteMode))
}
