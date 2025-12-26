// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveSchedulerJobsFromSchedulerJobGroup removes SchedulerJobsFromSchedulerJobGroup
func (cli *ZSClient) RemoveSchedulerJobsFromSchedulerJobGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobgroups/{schedulerJobGroupUuid}/job", uuid, string(deleteMode))
}
