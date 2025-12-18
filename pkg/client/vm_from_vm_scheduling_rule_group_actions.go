// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachVmFromVmSchedulingRuleGroup 操作VmFromVmSchedulingRuleGroup
func (cli *ZSClient) DetachVmFromVmSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vmSchedulingRuleGroup/{vmGroupUuid}/vmInstance/", uuid, string(deleteMode))
}

