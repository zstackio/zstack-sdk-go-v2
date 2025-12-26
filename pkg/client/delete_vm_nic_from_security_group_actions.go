// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmNicFromSecurityGroup deletes VmNicFromSecurityGroup
func (cli *ZSClient) DeleteVmNicFromSecurityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups/{securityGroupUuid}/vm-instances/nics", uuid, string(deleteMode))
}
