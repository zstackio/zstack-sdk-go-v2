// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteMonitorTriggerAction 删除MonitorTriggerAction
func (cli *ZSClient) DeleteMonitorTriggerAction(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/trigger-actions/{uuid}", uuid, string(deleteMode))
}

