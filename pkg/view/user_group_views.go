// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserGroupInventoryView UserGroup
type UserGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
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

