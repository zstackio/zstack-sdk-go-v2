// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveVmSchedulingRule 操作RemoveVmSchedulingRule
func (cli *ZSClient) RemoveVmSchedulingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vmSchedulingRule/{uuid}", uuid, string(deleteMode))
}

