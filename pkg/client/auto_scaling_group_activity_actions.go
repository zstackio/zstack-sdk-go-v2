// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAutoScalingGroupActivity queries AutoScalingGroupActivity list
func (cli *ZSClient) QueryAutoScalingGroupActivity(params *param.QueryParam) ([]view.AutoScalingGroupActivityInventoryView, error) {
	var resp []view.AutoScalingGroupActivityInventoryView
	return resp, cli.List("v1/autoscaling/groups/activities", params, &resp)
}
