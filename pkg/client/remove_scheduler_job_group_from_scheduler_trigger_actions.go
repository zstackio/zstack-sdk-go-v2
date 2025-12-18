// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveSchedulerJobGroupFromSchedulerTrigger 操作RemoveSchedulerJobGroupFromSchedulerTrigger
func (cli *ZSClient) RemoveSchedulerJobGroupFromSchedulerTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobgroups/{schedulerJobGroupUuid}/scheduler/triggers/{schedulerTriggerUuid}", uuid, string(deleteMode))
}

