// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserGroupUserRefInventoryView UserGroupUserRef
type UserGroupUserRefInventoryView struct {
	UserUuid string `json:"userUuid,omitempty"`
	GroupUuid string `json:"groupUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

