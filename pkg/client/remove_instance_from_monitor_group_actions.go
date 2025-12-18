// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveInstanceFromMonitorGroup 操作RemoveInstanceFromMonitorGroup
func (cli *ZSClient) RemoveInstanceFromMonitorGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitorgroups/{groupUuid}/actions/{instanceUuid}", uuid, string(deleteMode))
}

