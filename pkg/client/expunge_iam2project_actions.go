// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// ExpungeIAM2Project operates on IAM2Project
func (cli *ZSClient) ExpungeIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{uuid}/actions", uuid, string(deleteMode))
}
