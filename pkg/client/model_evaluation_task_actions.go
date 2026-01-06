// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteModelEvaluationTask deletes ModelEvaluationTask
func (cli *ZSClient) DeleteModelEvaluationTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/model-evaluation-tasks/{uuid}", uuid, string(deleteMode))
}
// QueryModelEvaluationTask queries ModelEvaluationTask list
func (cli *ZSClient) QueryModelEvaluationTask(params *param.QueryParam) ([]view.ModelEvaluationTaskInventoryView, error) {
	var resp []view.ModelEvaluationTaskInventoryView
	return resp, cli.List("v1/model-evaluation-tasks", params, &resp)
}
// UpdateModelEvaluationTask updates ModelEvaluationTask
func (cli *ZSClient) UpdateModelEvaluationTask(uuid string, params param.UpdateModelEvaluationTaskParam) (*view.ModelEvaluationTaskInventoryView, error) {
	var resp view.UpdateModelEvaluationTaskEventView
	if err := cli.Put("v1/model-evaluation-tasks/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
