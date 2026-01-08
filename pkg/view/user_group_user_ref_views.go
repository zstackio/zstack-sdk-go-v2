// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserGroupUserRefInventoryView UserGroupUserRef
type UserGroupUserRefInventoryView struct {
	UserUuid   string    `json:"userUuid,omitempty"`
	GroupUuid  string    `json:"groupUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}
