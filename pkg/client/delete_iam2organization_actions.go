// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteIAM2Organization deletes IAM2Organization
func (cli *ZSClient) DeleteIAM2Organization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/organizations/{uuid}", uuid, string(deleteMode))
}
