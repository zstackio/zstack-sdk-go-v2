// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveIAM2ProjectLoginExpired removes IAM2ProjectLoginExpired
func (cli *ZSClient) RemoveIAM2ProjectLoginExpired(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/remove/login-expired/{uuid}/actions", uuid, string(deleteMode))
}
