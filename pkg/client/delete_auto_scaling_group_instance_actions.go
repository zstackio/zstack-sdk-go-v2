// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAutoScalingGroupInstance deletes AutoScalingGroupInstance
func (cli *ZSClient) DeleteAutoScalingGroupInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/groups/instances/{instanceUuid}", uuid, string(deleteMode))
}
