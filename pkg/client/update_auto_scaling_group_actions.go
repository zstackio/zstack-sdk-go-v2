// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAutoScalingGroup updates AutoScalingGroup
func (cli *ZSClient) UpdateAutoScalingGroup(uuid string, params param.UpdateAutoScalingGroupParam) (*view.UpdateAutoScalingGroupEventView, error) {
	resp := view.UpdateAutoScalingGroupEventView{}
	if err := cli.Put("v1/autoscaling/groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
