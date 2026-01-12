// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HostHaStateInventoryView HostHaState
type HostHaStateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	State *string `json:"state,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
}

