// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSchedulerTrigger deletes SchedulerTrigger
func (cli *ZSClient) DeleteSchedulerTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/triggers/{uuid}", uuid, string(deleteMode))
}
