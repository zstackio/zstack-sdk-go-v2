// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AutoScalingTemplateGroupRefInventoryView AutoScalingTemplateGroupRef
type AutoScalingTemplateGroupRefInventoryView struct {
	TemplateUuid string `json:"templateUuid,omitempty"`
	GroupUuid string `json:"groupUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

