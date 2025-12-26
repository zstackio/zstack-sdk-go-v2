// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteMonitorGroup deletes MonitorGroup
func (cli *ZSClient) DeleteMonitorGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitorgroups/{uuid}", uuid, string(deleteMode))
}
