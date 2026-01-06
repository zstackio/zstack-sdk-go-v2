// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryModelServiceInstanceGroup queries ModelServiceInstanceGroup list
func (cli *ZSClient) QueryModelServiceInstanceGroup(params *param.QueryParam) ([]view.ModelServiceInstanceGroupInventoryView, error) {
	var resp []view.ModelServiceInstanceGroupInventoryView
	return resp, cli.List("v1/ai/model-services/instances/groups/", params, &resp)
}
// UpdateModelServiceInstanceGroup updates ModelServiceInstanceGroup
func (cli *ZSClient) UpdateModelServiceInstanceGroup(uuid string, params param.UpdateModelServiceInstanceGroupParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.UpdateModelServiceInstanceGroupEventView
	if err := cli.Put("v1/model-service-instance-groups/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteModelServiceInstanceGroup deletes ModelServiceInstanceGroup
func (cli *ZSClient) DeleteModelServiceInstanceGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-services/instances/groups/{uuid}", uuid, string(deleteMode))
}
