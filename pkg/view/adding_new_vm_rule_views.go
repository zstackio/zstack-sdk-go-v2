// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AddingNewVmRuleInventoryView AddingNewVmRule
type AddingNewVmRuleInventoryView struct {
	rest string `json:"adjustmentType,omitempty"`
	rest int `json:"adjustmentValue,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"description,omitempty"`
	rest int64 `json:"cooldown,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest []string `json:"systemTags,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"scalingGroupUuid,omitempty"`
	rest []AutoScalingRuleTriggerInventoryView `json:"ruleTriggers,omitempty"`
}

