// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// QuotaInventoryView Quota
type QuotaInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	IdentityUuid string `json:"identityUuid,omitempty"`
	IdentityType string `json:"identityType,omitempty"`
	Value int64 `json:"value,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
}

