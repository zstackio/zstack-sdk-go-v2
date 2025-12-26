// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteIAM2ProjectTemplate deletes IAM2ProjectTemplate
func (cli *ZSClient) DeleteIAM2ProjectTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/templates/{uuid}", uuid, string(deleteMode))
}
