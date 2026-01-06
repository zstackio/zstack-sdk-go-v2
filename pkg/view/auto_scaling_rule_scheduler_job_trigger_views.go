// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AutoScalingRuleSchedulerJobTriggerInventoryView AutoScalingRuleSchedulerJobTrigger
type AutoScalingRuleSchedulerJobTriggerInventoryView struct {
	SchedulerJobUuid string `json:"schedulerJobUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Type string `json:"type,omitempty"`
	RuleUuid string `json:"ruleUuid,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

