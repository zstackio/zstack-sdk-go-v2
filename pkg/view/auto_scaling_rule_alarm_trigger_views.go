// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AutoScalingRuleAlarmTriggerInventoryView AutoScalingRuleAlarmTrigger
type AutoScalingRuleAlarmTriggerInventoryView struct {
	BaseInfoView
	BaseTimeView
	AlarmUuid string `json:"alarmUuid,omitempty"`
	Type      string `json:"type,omitempty"`
	RuleUuid  string `json:"ruleUuid,omitempty"`
	State     string `json:"state,omitempty"`
}
