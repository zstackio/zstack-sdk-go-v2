// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmSchedulingRuleGroup deletes VmSchedulingRuleGroup
func (cli *ZSClient) DeleteVmSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vmSchedulingRuleGroup/{uuid}", uuid, string(deleteMode))
}
