// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAutoScalingRuleTrigger deletes AutoScalingRuleTrigger
func (cli *ZSClient) DeleteAutoScalingRuleTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/groups/rules/triggers/{uuid}", uuid, string(deleteMode))
}
