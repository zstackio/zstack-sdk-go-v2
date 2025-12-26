// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveSchedulerJobFromSchedulerTrigger removes SchedulerJobFromSchedulerTrigger
func (cli *ZSClient) RemoveSchedulerJobFromSchedulerTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobs/{schedulerJobUuid}/scheduler/triggers/{schedulerTriggerUuid}", uuid, string(deleteMode))
}
