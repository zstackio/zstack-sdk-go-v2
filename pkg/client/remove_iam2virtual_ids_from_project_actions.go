// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveIAM2VirtualIDsFromProject removes IAM2VirtualIDsFromProject
func (cli *ZSClient) RemoveIAM2VirtualIDsFromProject(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{projectUuid}/virtual-ids", uuid, string(deleteMode))
}
