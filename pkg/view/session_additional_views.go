// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SessionInventoryView Session
type SessionInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	UserUuid string `json:"userUuid,omitempty"`
	UserType string `json:"userType,omitempty"`
	ExpiredDate ZStackTime `json:"expiredDate,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
}

