// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AutoScalingGroupActivityInventoryView AutoScalingGroupActivity
type AutoScalingGroupActivityInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	ScalingGroupUuid string `json:"scalingGroupUuid,omitempty"`
	ActivityAction string `json:"activityAction,omitempty"`
	InstanceUuids string `json:"instanceUuids,omitempty"`
	ScalingGroupRuleUuid string `json:"scalingGroupRuleUuid,omitempty"`
	Cause string `json:"cause,omitempty"`
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
	ActivityActionResultMessage string `json:"activityActionResultMessage,omitempty"`
	EndDate time.Time `json:"endDate,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

