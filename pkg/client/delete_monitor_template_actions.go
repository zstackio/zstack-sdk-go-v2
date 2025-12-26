// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteMonitorTemplate deletes MonitorTemplate
func (cli *ZSClient) DeleteMonitorTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitortemplates/{uuid}", uuid, string(deleteMode))
}
