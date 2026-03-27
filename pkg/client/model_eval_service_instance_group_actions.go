// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryModelEvalServiceInstanceGroup queries ModelEvalServiceInstanceGroup list
func (cli *ZSClient) QueryModelEvalServiceInstanceGroup(ctx context.Context, params *param.QueryParam) ([]view.ModelServiceInstanceGroupInventoryView, error) {
	var resp []view.ModelServiceInstanceGroupInventoryView
	return resp, cli.List(ctx, "v1/ai/model-eval-services/instances/groups/", params, &resp)
}

func (cli *ZSClient) GetModelEvalServiceInstanceGroup(ctx context.Context, uuid string) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.ModelServiceInstanceGroupInventoryView
	if err := cli.Get(ctx, "v1/ai/model-eval-services/instances/groups/", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageModelEvalServiceInstanceGroup Pagination
func (cli *ZSClient) PageModelEvalServiceInstanceGroup(ctx context.Context, params *param.QueryParam) ([]view.ModelServiceInstanceGroupInventoryView, int, error) {
	var modelEvalServiceInstanceGroups []view.ModelServiceInstanceGroupInventoryView
	total, err := cli.Page(ctx, "v1/ai/model-eval-services/instances/groups/", params, &modelEvalServiceInstanceGroups)
	return modelEvalServiceInstanceGroups, total, err
}
