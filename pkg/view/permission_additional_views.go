// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PermissionView Permission
type PermissionView struct {
	Allow bool `json:"allow,omitempty"`
}
