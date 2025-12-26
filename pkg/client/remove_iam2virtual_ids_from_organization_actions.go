// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveIAM2VirtualIDsFromOrganization removes IAM2VirtualIDsFromOrganization
func (cli *ZSClient) RemoveIAM2VirtualIDsFromOrganization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/organizations/{organizationUuid}/virtual-ids", uuid, string(deleteMode))
}
