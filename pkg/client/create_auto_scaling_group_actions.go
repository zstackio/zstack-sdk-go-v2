// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAutoScalingGroup creates AutoScalingGroup
func (cli *ZSClient) CreateAutoScalingGroup(params param.CreateAutoScalingGroupParam) (*view.CreateAutoScalingGroupEventView, error) {
	resp := view.CreateAutoScalingGroupEventView{}
	if err := cli.Post("v1/autoscaling/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
