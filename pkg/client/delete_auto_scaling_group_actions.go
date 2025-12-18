// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAutoScalingGroup deletes AutoScalingGroup
func (cli *ZSClient) DeleteAutoScalingGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/groups/{uuid}", uuid, string(deleteMode))
}
