// Copyright (c) ZStack.io, Inc.

package view

import "time"

// QuotaInventoryView Quota
type QuotaInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"identityUuid,omitempty"`
	rest string `json:"identityType,omitempty"`
	rest int64 `json:"value,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
}

