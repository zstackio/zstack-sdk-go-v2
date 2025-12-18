// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExecuteAutoScalingRule 操作ExecuteAutoScalingRule
func (cli *ZSClient) ExecuteAutoScalingRule(uuid string, params param.ExecuteAutoScalingRuleParam) (*view.ExecuteAutoScalingRuleEventView, error) {
	resp := view.ExecuteAutoScalingRuleEventView{}
	if err := cli.Put("v1/autoscaling/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

