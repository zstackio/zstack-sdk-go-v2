// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveVmSchedulingRule removes VmSchedulingRule
func (cli *ZSClient) RemoveVmSchedulingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vmSchedulingRule/{uuid}", uuid, string(deleteMode))
}
