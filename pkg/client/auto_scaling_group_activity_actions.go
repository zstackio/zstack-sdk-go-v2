// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAutoScalingGroupActivity queries AutoScalingGroupActivity list
func (cli *ZSClient) QueryAutoScalingGroupActivity(params *param.QueryParam) ([]view.AutoScalingGroupActivityInventoryView, error) {
	var resp []view.AutoScalingGroupActivityInventoryView
	return resp, cli.List("v1/autoscaling/groups/activities", params, &resp)
}

func (cli *ZSClient) GetAutoScalingGroupActivity(uuid string) (*view.AutoScalingGroupActivityInventoryView, error) {
	var resp view.AutoScalingGroupActivityInventoryView
	if err := cli.Get("v1/autoscaling/groups/activities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAutoScalingGroupActivity Pagination
func (cli *ZSClient) PageAutoScalingGroupActivity(params *param.QueryParam) ([]view.AutoScalingGroupActivityInventoryView, int, error) {
	var autoScalingGroupActivities []view.AutoScalingGroupActivityInventoryView
	total, err := cli.Page("v1/autoscaling/groups/activities", params, &autoScalingGroupActivities)
	return autoScalingGroupActivities, total, err
}
