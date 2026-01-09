// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAutoScalingRuleAlarmTrigger creates AutoScalingRuleAlarmTrigger
func (cli *ZSClient) CreateAutoScalingRuleAlarmTrigger(params param.CreateAutoScalingRuleAlarmTriggerParam) (*view.AutoScalingRuleTriggerInventoryView, error) {
	var resp view.CreateAutoScalingRuleTriggerEventView
	if err := cli.Post("v1/zwatch/alarms/{alarmUuid}/autoscaling/rules/{ruleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
