// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteModelEvaluationTask deletes ModelEvaluationTask
func (cli *ZSClient) DeleteModelEvaluationTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/model-evaluation-tasks/{uuid}", uuid, string(deleteMode))
}
