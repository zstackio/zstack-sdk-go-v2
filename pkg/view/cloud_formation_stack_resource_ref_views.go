// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CloudFormationStackResourceRefInventoryView CloudFormationStackResourceRef
type CloudFormationStackResourceRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"stackUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceName,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest bool `json:"reserve,omitempty"`
	rest int `json:"round,omitempty"`
}

