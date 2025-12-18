// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAutoScalingRule deletes AutoScalingRule
func (cli *ZSClient) DeleteAutoScalingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/rules/{uuid}", uuid, string(deleteMode))
}
