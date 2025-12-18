// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BillingInventoryView Billing
type BillingInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"billingType,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceName,omitempty"`
	rest float64 `json:"spending,omitempty"`
	rest int64 `json:"startTime,omitempty"`
	rest int64 `json:"endTime,omitempty"`
	rest string `json:"hypervisorType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

