// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAutoScalingRuleAlarmTrigger creates AutoScalingRuleAlarmTrigger
func (cli *ZSClient) CreateAutoScalingRuleAlarmTrigger(alarmUuid string, ruleUuid string, params param.CreateAutoScalingRuleAlarmTriggerParam) (*view.AutoScalingRuleTriggerInventoryView, error) {
	resp := view.AutoScalingRuleTriggerInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/zwatch/alarms/%s/autoscaling/rules/%s", alarmUuid, ruleUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
