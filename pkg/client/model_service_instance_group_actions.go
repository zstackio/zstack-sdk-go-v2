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
// UpdateModelServiceInstanceGroup updates ModelServiceInstanceGroup
func (cli *ZSClient) UpdateModelServiceInstanceGroup(uuid string, params param.UpdateModelServiceInstanceGroupParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.UpdateModelServiceInstanceGroupEventView
	err := cli.PutWithSpec("v1/model-service-instance-groups", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteModelServiceInstanceGroup deletes ModelServiceInstanceGroup
func (cli *ZSClient) DeleteModelServiceInstanceGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/ai/model-services/instances/groups", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
