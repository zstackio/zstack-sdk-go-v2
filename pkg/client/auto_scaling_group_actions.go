// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteAutoScalingGroup deletes AutoScalingGroup
func (cli *ZSClient) DeleteAutoScalingGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/groups", uuid, string(deleteMode))
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

func (cli *ZSClient) GetAutoScalingGroup(uuid string) (*view.AutoScalingGroupInventoryView, error) {
	var resp view.AutoScalingGroupInventoryView
	if err := cli.Get("v1/autoscaling/groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
