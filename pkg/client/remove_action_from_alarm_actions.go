// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveActionFromAlarm removes ActionFromAlarm
func (cli *ZSClient) RemoveActionFromAlarm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms/{alarmUuid}/actions/{actionUuid}", uuid, string(deleteMode))
}
