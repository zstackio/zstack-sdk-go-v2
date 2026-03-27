// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteModelEvaluationTask deletes ModelEvaluationTask
func (cli *ZSClient) DeleteModelEvaluationTask(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/model-evaluation-tasks", uuid, string(deleteMode))
}
// QueryModelEvaluationTask queries ModelEvaluationTask list
func (cli *ZSClient) QueryModelEvaluationTask(ctx context.Context, params *param.QueryParam) ([]view.ModelEvaluationTaskInventoryView, error) {
	var resp []view.ModelEvaluationTaskInventoryView
	return resp, cli.List(ctx, "v1/model-evaluation-tasks", params, &resp)
}

func (cli *ZSClient) GetModelEvaluationTask(ctx context.Context, uuid string) (*view.ModelEvaluationTaskInventoryView, error) {
	var resp view.ModelEvaluationTaskInventoryView
	if err := cli.Get(ctx, "v1/model-evaluation-tasks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageModelEvaluationTask Pagination
func (cli *ZSClient) PageModelEvaluationTask(ctx context.Context, params *param.QueryParam) ([]view.ModelEvaluationTaskInventoryView, int, error) {
	var modelEvaluationTasks []view.ModelEvaluationTaskInventoryView
	total, err := cli.Page(ctx, "v1/model-evaluation-tasks", params, &modelEvaluationTasks)
	return modelEvaluationTasks, total, err
}
// UpdateModelEvaluationTask updates ModelEvaluationTask
func (cli *ZSClient) UpdateModelEvaluationTask(ctx context.Context, uuid string, params param.UpdateModelEvaluationTaskParam) (*view.ModelEvaluationTaskInventoryView, error) {
	resp := view.ModelEvaluationTaskInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/model-evaluation-tasks", uuid, "", map[string]interface{}{
		"updateModelEvaluationTask": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
