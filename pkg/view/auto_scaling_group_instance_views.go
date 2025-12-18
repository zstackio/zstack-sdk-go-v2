// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AutoScalingGroupInstanceInventoryView AutoScalingGroupInstance
type AutoScalingGroupInstanceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"instanceUuid,omitempty"`
	rest string `json:"scalingGroupUuid,omitempty"`
	rest string `json:"templateUuid,omitempty"`
	rest string `json:"scalingGroupActivityUuid,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"healthStatus,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"protectionStrategy,omitempty"`
}

