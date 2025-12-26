// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSchedulerJobGroup deletes SchedulerJobGroup
func (cli *ZSClient) DeleteSchedulerJobGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobgroups/{uuid}", uuid, string(deleteMode))
}
