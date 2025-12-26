// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveIAM2VirtualIDGroupFromProjects removes IAM2VirtualIDGroupFromProjects
func (cli *ZSClient) RemoveIAM2VirtualIDGroupFromProjects(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups", uuid, string(deleteMode))
}
