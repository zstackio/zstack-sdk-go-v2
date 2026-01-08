// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// GuestOsCategoryInventoryView GuestOsCategory
type GuestOsCategoryInventoryView struct {
	Uuid      string `json:"uuid,omitempty"`
	Platform  string `json:"platform,omitempty"`
	Name      string `json:"name,omitempty"`
	OsRelease string `json:"osRelease,omitempty"`
	Version   string `json:"version,omitempty"`
}
