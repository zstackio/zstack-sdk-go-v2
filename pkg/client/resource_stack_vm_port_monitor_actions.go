// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteResourceStackVmPortMonitor 删除ResourceStackVmPortMonitor
func (cli *ZSClient) DeleteResourceStackVmPortMonitor(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cloudformation/stack/monitor/delvm", uuid, string(deleteMode))
}

