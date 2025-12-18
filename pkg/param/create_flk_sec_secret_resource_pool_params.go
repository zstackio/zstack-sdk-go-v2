// Copyright (c) ZStack.io, Inc.

package param

// CreateFlkSecSecretResourcePoolDetailParam CreateFlkSecSecretResourcePool detail param
type CreateFlkSecSecretResourcePoolDetailParam struct {
	TestResult string `json:"testResult,omitempty"`
	ActivatedToken string `json:"activatedToken,omitempty"`
	ProtectToken string `json:"protectToken,omitempty"`
	HmacToken string `json:"hmacToken,omitempty"`
	UkeyType string `json:"ukeyType,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	Ability string `json:"ability,omitempty"`
	Type string `json:"type" validate:"required"`
	HeartbeatInterval int `json:"heartbeatInterval" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFlkSecSecretResourcePoolParam CreateFlkSecSecretResourcePool request param
type CreateFlkSecSecretResourcePoolParam struct {
	BaseParam
	Params CreateFlkSecSecretResourcePoolDetailParam `json:"params"`
}
