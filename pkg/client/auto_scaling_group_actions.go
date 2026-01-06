// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteAutoScalingGroup deletes AutoScalingGroup
func (cli *ZSClient) DeleteAutoScalingGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/groups/{uuid}", uuid, string(deleteMode))
}
// CreateAutoScalingGroup creates AutoScalingGroup
func (cli *ZSClient) CreateAutoScalingGroup(params param.CreateAutoScalingGroupParam) (*view.AutoScalingGroupInventoryView, error) {
	var resp view.CreateAutoScalingGroupEventView
	if err := cli.Post("v1/autoscaling/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateAutoScalingGroup updates AutoScalingGroup
func (cli *ZSClient) UpdateAutoScalingGroup(uuid string, params param.UpdateAutoScalingGroupParam) (*view.AutoScalingGroupInventoryView, error) {
	var resp view.UpdateAutoScalingGroupEventView
	if err := cli.Put("v1/autoscaling/groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAutoScalingGroup queries AutoScalingGroup list
func (cli *ZSClient) QueryAutoScalingGroup(params *param.QueryParam) ([]view.AutoScalingGroupInventoryView, error) {
	var resp []view.AutoScalingGroupInventoryView
	return resp, cli.List("v1/autoscaling/groups", params, &resp)
}
