// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteModelEvaluationTask deletes ModelEvaluationTask
func (cli *ZSClient) DeleteModelEvaluationTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/model-evaluation-tasks", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryModelEvaluationTask queries ModelEvaluationTask list
func (cli *ZSClient) QueryModelEvaluationTask(params *param.QueryParam) ([]view.ModelEvaluationTaskInventoryView, error) {
	var resp []view.ModelEvaluationTaskInventoryView
	return resp, cli.List("v1/model-evaluation-tasks", params, &resp)
}

func (cli *ZSClient) GetModelEvaluationTask(uuid string) (*view.ModelEvaluationTaskInventoryView, error) {
	var resp view.ModelEvaluationTaskInventoryView
	if err := cli.Get("v1/model-evaluation-tasks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateModelEvaluationTask updates ModelEvaluationTask
func (cli *ZSClient) UpdateModelEvaluationTask(uuid string, params param.UpdateModelEvaluationTaskParam) (*view.ModelEvaluationTaskInventoryView, error) {
	var resp view.UpdateModelEvaluationTaskEventView
	err := cli.PutWithSpec("v1/model-evaluation-tasks", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
