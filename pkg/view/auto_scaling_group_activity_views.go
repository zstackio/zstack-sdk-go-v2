// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AutoScalingGroupActivityInventoryView AutoScalingGroupActivity
type AutoScalingGroupActivityInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"scalingGroupUuid,omitempty"`
	rest string `json:"activityAction,omitempty"`
	rest string `json:"instanceUuids,omitempty"`
	rest string `json:"scalingGroupRuleUuid,omitempty"`
	rest string `json:"cause,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"activityActionResultMessage,omitempty"`
	rest time.Time `json:"endDate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

