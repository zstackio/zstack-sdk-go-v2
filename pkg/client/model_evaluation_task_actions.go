// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateModelEvaluationTask 更新ModelEvaluationTask
func (cli *ZSClient) UpdateModelEvaluationTask(uuid string, params param.UpdateModelEvaluationTaskParam) (*view.UpdateModelEvaluationTaskEventView, error) {
	resp := view.UpdateModelEvaluationTaskEventView{}
	if err := cli.Put("v1/model-evaluation-tasks/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

