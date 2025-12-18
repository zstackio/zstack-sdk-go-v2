// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAutoScalingRuleSchedulerJobTrigger 创建AutoScalingRuleSchedulerJobTrigger
func (cli *ZSClient) CreateAutoScalingRuleSchedulerJobTrigger(params param.CreateAutoScalingRuleSchedulerJobTriggerParam) (*view.CreateAutoScalingRuleTriggerEventView, error) {
	resp := view.CreateAutoScalingRuleTriggerEventView{}
	if err := cli.Post("v1/scheduler/jobs/{schedulerJobUuid}/autoscaling/rules/{ruleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

