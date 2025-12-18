// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CloudFormationStackEventInventoryView CloudFormationStackEvent
type CloudFormationStackEventInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"action,omitempty"`
	rest string `json:"content,omitempty"`
	rest string `json:"resourceName,omitempty"`
	rest string `json:"actionStatus,omitempty"`
	rest string `json:"stackUuid,omitempty"`
	rest string `json:"duration,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

