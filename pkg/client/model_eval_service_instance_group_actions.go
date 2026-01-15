// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryModelEvalServiceInstanceGroup queries ModelEvalServiceInstanceGroup list
func (cli *ZSClient) QueryModelEvalServiceInstanceGroup(params *param.QueryParam) ([]view.ModelServiceInstanceGroupInventoryView, error) {
	var resp []view.ModelServiceInstanceGroupInventoryView
	return resp, cli.List("v1/ai/model-eval-services/instances/groups/", params, &resp)
}

// PageModelEvalServiceInstanceGroup Pagination
func (cli *ZSClient) PageModelEvalServiceInstanceGroup(params *param.QueryParam) ([]view.ModelServiceInstanceGroupInventoryView, int, error) {
	var modelEvalServiceInstanceGroups []view.ModelServiceInstanceGroupInventoryView
	total, err := cli.Page("v1/ai/model-eval-services/instances/groups/", params, &modelEvalServiceInstanceGroups)
	return modelEvalServiceInstanceGroups, total, err
}
