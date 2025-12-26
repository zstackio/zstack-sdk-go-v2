// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteMonitorTriggerAction deletes MonitorTriggerAction
func (cli *ZSClient) DeleteMonitorTriggerAction(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/trigger-actions/{uuid}", uuid, string(deleteMode))
}
