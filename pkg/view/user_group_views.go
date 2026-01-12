// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserGroupInventoryView UserGroup
type UserGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccountUuid *string `json:"accountUuid,omitempty"`
	Description *string `json:"description,omitempty"`
}

// DeleteUserGroupEventView DeleteUserGroupEvent
type DeleteUserGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateUserGroupEventView CreateUserGroupEvent
type CreateUserGroupEventView struct {
	Inventory UserGroupInventoryView `json:"inventory,omitempty"`
}

// QueryUserGroupView QueryUserGroup
type QueryUserGroupView struct {
	Inventories []UserGroupInventoryView `json:"inventories,omitempty"`
}

// UpdateUserGroupEventView UpdateUserGroupEvent
type UpdateUserGroupEventView struct {
	Inventory UserGroupInventoryView `json:"inventory,omitempty"`
}

