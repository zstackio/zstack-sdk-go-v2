// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVmNicFromSecurityGroup 删除VmNicFromSecurityGroup
func (cli *ZSClient) DeleteVmNicFromSecurityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups/{securityGroupUuid}/vm-instances/nics", uuid, string(deleteMode))
}

