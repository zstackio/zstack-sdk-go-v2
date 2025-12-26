// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachVmFromVmSchedulingRuleGroup operates on VmFromVmSchedulingRuleGroup
func (cli *ZSClient) DetachVmFromVmSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vmSchedulingRuleGroup/{vmGroupUuid}/vmInstance/", uuid, string(deleteMode))
}
