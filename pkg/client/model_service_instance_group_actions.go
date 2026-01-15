// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryModelServiceInstanceGroup queries ModelServiceInstanceGroup list
func (cli *ZSClient) QueryModelServiceInstanceGroup(params *param.QueryParam) ([]view.ModelServiceInstanceGroupInventoryView, error) {
	var resp []view.ModelServiceInstanceGroupInventoryView
	return resp, cli.List("v1/ai/model-services/instances/groups/", params, &resp)
}

func (cli *ZSClient) GetModelServiceInstanceGroup(uuid string) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.ModelServiceInstanceGroupInventoryView
	if err := cli.Get("v1/ai/model-services/instances/groups/", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageModelServiceInstanceGroup Pagination
func (cli *ZSClient) PageModelServiceInstanceGroup(params *param.QueryParam) ([]view.ModelServiceInstanceGroupInventoryView, int, error) {
	var modelServiceInstanceGroups []view.ModelServiceInstanceGroupInventoryView
	total, err := cli.Page("v1/ai/model-services/instances/groups/", params, &modelServiceInstanceGroups)
	return modelServiceInstanceGroups, total, err
}
// UpdateModelServiceInstanceGroup updates ModelServiceInstanceGroup
func (cli *ZSClient) UpdateModelServiceInstanceGroup(uuid string, params param.UpdateModelServiceInstanceGroupParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	resp := view.ModelServiceInstanceGroupInventoryView{}
	if err := cli.Put("v1/model-service-instance-groups", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteModelServiceInstanceGroup deletes ModelServiceInstanceGroup
func (cli *ZSClient) DeleteModelServiceInstanceGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-services/instances/groups", uuid, string(deleteMode))
}
