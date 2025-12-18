// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteModelEvaluationTasks 删除ModelEvaluationTasks
func (cli *ZSClient) DeleteModelEvaluationTasks(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/model-evaluation-tasks", uuid, string(deleteMode))
}

