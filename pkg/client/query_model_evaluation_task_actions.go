// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryModelEvaluationTask queries ModelEvaluationTask list
func (cli *ZSClient) QueryModelEvaluationTask(params *param.QueryParam) ([]view.ModelEvaluationTaskInventoryView, error) {
	var resp []view.ModelEvaluationTaskInventoryView
	return resp, cli.List("v1/model-evaluation-tasks", params, &resp)
}
