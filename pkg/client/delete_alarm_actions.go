// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAlarm deletes Alarm
func (cli *ZSClient) DeleteAlarm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms/{uuid}", uuid, string(deleteMode))
}
