// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteHostSchedulingRuleGroup deletes HostSchedulingRuleGroup
func (cli *ZSClient) DeleteHostSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hostSchedulingRuleGroup/{uuid}", uuid, string(deleteMode))
}
