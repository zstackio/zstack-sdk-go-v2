// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SSOTokenInventoryView SSOToken
type SSOTokenInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ClientUuid string `json:"clientUuid,omitempty"`
	UserUuid string `json:"userUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

