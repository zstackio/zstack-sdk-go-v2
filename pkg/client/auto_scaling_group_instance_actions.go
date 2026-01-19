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
	resp := view.AutoScalingGroupInstanceInventoryView{}
	if err := cli.Put("v1/autoscaling/groups/instances", instanceUuid, map[string]interface{}{
		"updateAutoScalingGroupInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAutoScalingGroupInstance deletes AutoScalingGroupInstance
func (cli *ZSClient) DeleteAutoScalingGroupInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/groups/instances", uuid, string(deleteMode))
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

// PageAutoScalingGroupInstance Pagination
func (cli *ZSClient) PageAutoScalingGroupInstance(params *param.QueryParam) ([]view.AutoScalingGroupInstanceInventoryView, int, error) {
	var autoScalingGroupInstances []view.AutoScalingGroupInstanceInventoryView
	total, err := cli.Page("v1/autoscaling/groups/instances", params, &autoScalingGroupInstances)
	return autoScalingGroupInstances, total, err
}
