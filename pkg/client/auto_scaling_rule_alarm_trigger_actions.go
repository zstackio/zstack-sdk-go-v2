// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAutoScalingRuleAlarmTrigger 创建AutoScalingRuleAlarmTrigger
func (cli *ZSClient) CreateAutoScalingRuleAlarmTrigger(params param.CreateAutoScalingRuleAlarmTriggerParam) (*view.CreateAutoScalingRuleTriggerEventView, error) {
	resp := view.CreateAutoScalingRuleTriggerEventView{}
	if err := cli.Post("v1/zwatch/alarms/{alarmUuid}/autoscaling/rules/{ruleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

