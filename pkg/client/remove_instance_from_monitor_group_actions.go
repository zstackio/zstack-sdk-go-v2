// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveInstanceFromMonitorGroup removes InstanceFromMonitorGroup
func (cli *ZSClient) RemoveInstanceFromMonitorGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitorgroups/{groupUuid}/actions/{instanceUuid}", uuid, string(deleteMode))
}
