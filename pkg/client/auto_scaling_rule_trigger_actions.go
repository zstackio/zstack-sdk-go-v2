// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAutoScalingRuleTrigger 删除AutoScalingRuleTrigger
func (cli *ZSClient) DeleteAutoScalingRuleTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/groups/rules/triggers/{uuid}", uuid, string(deleteMode))
}

