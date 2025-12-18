// Copyright (c) ZStack.io, Inc.

package view

import "time"

// FiSecSecretResourcePoolInventoryView FiSecSecretResourcePool
type FiSecSecretResourcePoolInventoryView struct {
	rest string `json:"keyNum,omitempty"`
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

