// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HybridAccountInventoryView HybridAccount
type HybridAccountInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"userUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"akey,omitempty"`
	rest string `json:"hybridAccountId,omitempty"`
	rest string `json:"hybridUserId,omitempty"`
	rest string `json:"hybridUserName,omitempty"`
	rest string `json:"current,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

