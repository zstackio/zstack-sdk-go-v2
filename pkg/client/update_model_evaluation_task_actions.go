// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateModelEvaluationTask updates ModelEvaluationTask
func (cli *ZSClient) UpdateModelEvaluationTask(uuid string, params param.UpdateModelEvaluationTaskParam) (*view.UpdateModelEvaluationTaskEventView, error) {
	resp := view.UpdateModelEvaluationTaskEventView{}
	if err := cli.Put("v1/model-evaluation-tasks/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
