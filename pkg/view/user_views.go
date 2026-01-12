// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserInventoryView User
type UserInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccountUuid *string `json:"accountUuid,omitempty"`
	Description *string `json:"description,omitempty"`
}

// QueryUserView QueryUser
type QueryUserView struct {
	Inventories []UserInventoryView `json:"inventories,omitempty"`
}

// DeleteUserEventView DeleteUserEvent
type DeleteUserEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateUserEventView UpdateUserEvent
type UpdateUserEventView struct {
	Inventory UserInventoryView `json:"inventory,omitempty"`
}

// CreateUserEventView CreateUserEvent
type CreateUserEventView struct {
	Inventory UserInventoryView `json:"inventory,omitempty"`
}

