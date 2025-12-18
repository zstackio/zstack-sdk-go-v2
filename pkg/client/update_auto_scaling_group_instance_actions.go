// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAutoScalingGroupInstance updates AutoScalingGroupInstance
func (cli *ZSClient) UpdateAutoScalingGroupInstance(uuid string, params param.UpdateAutoScalingGroupInstanceParam) (*view.UpdateAutoScalingGroupInstanceEventView, error) {
	resp := view.UpdateAutoScalingGroupInstanceEventView{}
	if err := cli.Put("v1/autoscaling/groups/instances/{instanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
