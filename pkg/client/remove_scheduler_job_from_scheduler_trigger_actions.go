// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveSchedulerJobFromSchedulerTrigger 操作RemoveSchedulerJobFromSchedulerTrigger
func (cli *ZSClient) RemoveSchedulerJobFromSchedulerTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobs/{schedulerJobUuid}/scheduler/triggers/{schedulerTriggerUuid}", uuid, string(deleteMode))
}

