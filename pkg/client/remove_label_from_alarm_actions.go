// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveLabelFromAlarm removes LabelFromAlarm
func (cli *ZSClient) RemoveLabelFromAlarm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms/labels/{uuid}", uuid, string(deleteMode))
}
