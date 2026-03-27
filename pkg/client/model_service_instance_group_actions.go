// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryModelServiceInstanceGroup queries ModelServiceInstanceGroup list
func (cli *ZSClient) QueryModelServiceInstanceGroup(ctx context.Context, params *param.QueryParam) ([]view.ModelServiceInstanceGroupInventoryView, error) {
	var resp []view.ModelServiceInstanceGroupInventoryView
	return resp, cli.List(ctx, "v1/ai/model-services/instances/groups/", params, &resp)
}

func (cli *ZSClient) GetModelServiceInstanceGroup(ctx context.Context, uuid string) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.ModelServiceInstanceGroupInventoryView
	if err := cli.Get(ctx, "v1/ai/model-services/instances/groups/", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageModelServiceInstanceGroup Pagination
func (cli *ZSClient) PageModelServiceInstanceGroup(ctx context.Context, params *param.QueryParam) ([]view.ModelServiceInstanceGroupInventoryView, int, error) {
	var modelServiceInstanceGroups []view.ModelServiceInstanceGroupInventoryView
	total, err := cli.Page(ctx, "v1/ai/model-services/instances/groups/", params, &modelServiceInstanceGroups)
	return modelServiceInstanceGroups, total, err
}
// UpdateModelServiceInstanceGroup updates ModelServiceInstanceGroup
func (cli *ZSClient) UpdateModelServiceInstanceGroup(ctx context.Context, uuid string, params param.UpdateModelServiceInstanceGroupParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	resp := view.ModelServiceInstanceGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/model-service-instance-groups", uuid, "", map[string]interface{}{
		"updateModelServiceInstanceGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteModelServiceInstanceGroup deletes ModelServiceInstanceGroup
func (cli *ZSClient) DeleteModelServiceInstanceGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ai/model-services/instances/groups", uuid, string(deleteMode))
}
