// Copyright (c) ZStack.io, Inc.

package view

import "time"

// UserGroupUserRefInventoryView UserGroupUserRef
type UserGroupUserRefInventoryView struct {
	rest string `json:"userUuid,omitempty"`
	rest string `json:"groupUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

