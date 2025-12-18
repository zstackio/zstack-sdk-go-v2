// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AutoScalingTemplateGroupRefInventoryView AutoScalingTemplateGroupRef
type AutoScalingTemplateGroupRefInventoryView struct {
	rest string `json:"templateUuid,omitempty"`
	rest string `json:"groupUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

