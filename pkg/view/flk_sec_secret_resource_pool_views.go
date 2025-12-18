// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// FlkSecSecretResourcePoolInventoryView FlkSecSecretResourcePool
type FlkSecSecretResourcePoolInventoryView struct {
	EncryptResult string `json:"encryptResult,omitempty"`
	ActivatedToken string `json:"activatedToken,omitempty"`
	ProtectToken string `json:"protectToken,omitempty"`
	HmacToken string `json:"hmacToken,omitempty"`
	UkeyType string `json:"ukeyType,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Ability string `json:"ability,omitempty"`
}

