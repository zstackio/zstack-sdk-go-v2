// Copyright (c) ZStack.io, Inc.

package view

// GetIAM2VirtualIDAPIPermissionView GetIAM2VirtualIDPermission
type GetIAM2VirtualIDAPIPermissionView struct {
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	NoPermission bool `json:"noPermission,omitempty"`
	Success bool `json:"success,omitempty"`
}

