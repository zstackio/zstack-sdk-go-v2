// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveLabelFromAlarm 操作RemoveLabelFromAlarm
func (cli *ZSClient) RemoveLabelFromAlarm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms/labels/{uuid}", uuid, string(deleteMode))
}

