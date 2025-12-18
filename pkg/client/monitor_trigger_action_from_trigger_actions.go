// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachMonitorTriggerActionFromTrigger 操作MonitorTriggerActionFromTrigger
func (cli *ZSClient) DetachMonitorTriggerActionFromTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/triggers/{triggerUuid}/trigger-actions/{actionUuid}", uuid, string(deleteMode))
}

