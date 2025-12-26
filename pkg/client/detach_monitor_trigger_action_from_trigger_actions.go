// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachMonitorTriggerActionFromTrigger operates on MonitorTriggerActionFromTrigger
func (cli *ZSClient) DetachMonitorTriggerActionFromTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/triggers/{triggerUuid}/trigger-actions/{actionUuid}", uuid, string(deleteMode))
}
