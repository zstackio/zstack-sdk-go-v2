// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAutoScalingGroupInstance deletes AutoScalingGroupInstance
func (cli *ZSClient) DeleteAutoScalingGroupInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/groups/instances/{instanceUuid}", uuid, string(deleteMode))
}
