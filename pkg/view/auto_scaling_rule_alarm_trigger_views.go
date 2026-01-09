// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AutoScalingRuleAlarmTriggerInventoryView AutoScalingRuleAlarmTrigger
type AutoScalingRuleAlarmTriggerInventoryView struct {
	AlarmUuid *string `json:"alarmUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Type *string `json:"type,omitempty"`
	RuleUuid *string `json:"ruleUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

