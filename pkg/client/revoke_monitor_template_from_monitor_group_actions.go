// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RevokeMonitorTemplateFromMonitorGroup operates on RevokeMonitorTemplateFromMonitorGroup
func (cli *ZSClient) RevokeMonitorTemplateFromMonitorGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitortemplates/{templateUuid}/monitorgroups/{groupUuid}", uuid, string(deleteMode))
}
