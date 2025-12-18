// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVmSchedulingRuleGroup 删除VmSchedulingRuleGroup
func (cli *ZSClient) DeleteVmSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vmSchedulingRuleGroup/{uuid}", uuid, string(deleteMode))
}

