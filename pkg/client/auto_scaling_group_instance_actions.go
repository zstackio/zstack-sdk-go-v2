// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAutoScalingGroupInstance updates AutoScalingGroupInstance
func (cli *ZSClient) UpdateAutoScalingGroupInstance(instanceUuid string, params param.UpdateAutoScalingGroupInstanceParam) (*view.AutoScalingGroupInstanceInventoryView, error) {
	var resp view.UpdateAutoScalingGroupInstanceEventView
	err := cli.PutWithSpec("v1/autoscaling/groups/instances", fmt.Sprintf(\"%s/actions\", instanceUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAutoScalingGroupInstance deletes AutoScalingGroupInstance
func (cli *ZSClient) DeleteAutoScalingGroupInstance(instanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/autoscaling/groups/instances", fmt.Sprintf(\"%s\", instanceUuid), string(deleteMode))
}
// QueryAutoScalingGroupInstance queries AutoScalingGroupInstance list
func (cli *ZSClient) QueryAutoScalingGroupInstance(params *param.QueryParam) ([]view.AutoScalingGroupInstanceInventoryView, error) {
	var resp []view.AutoScalingGroupInstanceInventoryView
	return resp, cli.List("v1/autoscaling/groups/instances", params, &resp)
}

func (cli *ZSClient) GetAutoScalingGroupInstance(uuid string) (*view.AutoScalingGroupInstanceInventoryView, error) {
	var resp view.AutoScalingGroupInstanceInventoryView
	if err := cli.Get("v1/autoscaling/groups/instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
