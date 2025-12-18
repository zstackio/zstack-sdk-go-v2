// Copyright (c) ZStack.io, Inc.

package view

import "time"

// FlkSecSecretResourcePoolInventoryView FlkSecSecretResourcePool
type FlkSecSecretResourcePoolInventoryView struct {
	rest string `json:"encryptResult,omitempty"`
	rest string `json:"activatedToken,omitempty"`
	rest string `json:"protectToken,omitempty"`
	rest string `json:"hmacToken,omitempty"`
	rest string `json:"ukeyType,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"model,omitempty"`
	rest int `json:"heartbeatInterval,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"ability,omitempty"`
}

