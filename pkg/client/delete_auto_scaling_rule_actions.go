// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAutoScalingRule deletes AutoScalingRule
func (cli *ZSClient) DeleteAutoScalingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/rules/{uuid}", uuid, string(deleteMode))
}
