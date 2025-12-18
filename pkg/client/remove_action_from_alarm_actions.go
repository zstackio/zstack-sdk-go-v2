// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveActionFromAlarm 操作RemoveActionFromAlarm
func (cli *ZSClient) RemoveActionFromAlarm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms/{alarmUuid}/actions/{actionUuid}", uuid, string(deleteMode))
}

