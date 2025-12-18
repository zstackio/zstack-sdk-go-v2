// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AutoScalingTemplateInventoryView AutoScalingTemplate
type AutoScalingTemplateInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest []string `json:"systemTags,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

