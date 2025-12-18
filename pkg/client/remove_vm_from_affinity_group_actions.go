// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveVmFromAffinityGroup 操作RemoveVmFromAffinityGroup
func (cli *ZSClient) RemoveVmFromAffinityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/affinity-groups/{affinityGroupUuid}/vm-instances", uuid, string(deleteMode))
}

