// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AutoScalingGroupInventoryView AutoScalingGroup
type AutoScalingGroupInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"scalingResourceType,omitempty"`
	rest string `json:"state,omitempty"`
	rest int64 `json:"defaultCooldown,omitempty"`
	rest string `json:"description,omitempty"`
	rest int `json:"minResourceSize,omitempty"`
	rest int `json:"maxResourceSize,omitempty"`
	rest string `json:"removalPolicy,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []string `json:"attachedTemplates,omitempty"`
	rest []string `json:"systemTags,omitempty"`
}

