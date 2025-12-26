// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveVmFromAffinityGroup removes VmFromAffinityGroup
func (cli *ZSClient) RemoveVmFromAffinityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/affinity-groups/{affinityGroupUuid}/vm-instances", uuid, string(deleteMode))
}
