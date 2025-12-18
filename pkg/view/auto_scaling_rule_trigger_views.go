// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AutoScalingRuleTriggerInventoryView AutoScalingRuleTrigger
type AutoScalingRuleTriggerInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"ruleUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

