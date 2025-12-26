// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteMonitorTrigger deletes MonitorTrigger
func (cli *ZSClient) DeleteMonitorTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/triggers/{uuid}", uuid, string(deleteMode))
}
