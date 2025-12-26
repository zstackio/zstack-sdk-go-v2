// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachIAM2ProjectFromIAM2Organization operates on IAM2ProjectFromIAM2Organization
func (cli *ZSClient) DetachIAM2ProjectFromIAM2Organization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{projectUuid}/iam2/organizations", uuid, string(deleteMode))
}
