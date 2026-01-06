// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserInventoryView User
type UserInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
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

