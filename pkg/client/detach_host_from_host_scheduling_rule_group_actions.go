// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachHostFromHostSchedulingRuleGroup operates on HostFromHostSchedulingRuleGroup
func (cli *ZSClient) DetachHostFromHostSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hostSchedulingRuleGroup/{hostGroupUuid}/host", uuid, string(deleteMode))
}
