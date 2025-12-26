// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAutoScalingGroup deletes AutoScalingGroup
func (cli *ZSClient) DeleteAutoScalingGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/groups/{uuid}", uuid, string(deleteMode))
}
