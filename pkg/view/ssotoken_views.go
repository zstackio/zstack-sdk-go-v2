// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SSOTokenInventoryView SSOToken
type SSOTokenInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ClientUuid *string `json:"clientUuid,omitempty"`
	UserUuid *string `json:"userUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

