// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAutoScalingGroupRemovalInstanceRule 更新AutoScalingGroupRemovalInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupRemovalInstanceRule(uuid string, params param.UpdateAutoScalingGroupRemovalInstanceRuleParam) (*view.UpdateAutoScalingRuleEventView, error) {
	resp := view.UpdateAutoScalingRuleEventView{}
	if err := cli.Put("v1/autoscaling/rules/removal-instance/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

