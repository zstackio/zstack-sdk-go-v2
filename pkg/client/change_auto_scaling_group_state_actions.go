// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeAutoScalingGroupState changes AutoScalingGroupState
func (cli *ZSClient) ChangeAutoScalingGroupState(uuid string, params param.ChangeAutoScalingGroupStateParam) (*view.ChangeAutoScalingGroupStateEventView, error) {
	resp := view.ChangeAutoScalingGroupStateEventView{}
	if err := cli.Put("v1/autoscaling/groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
