// Copyright (c) ZStack.io, Inc.

package param

// CreateJitSecretResourcePoolDetailParam CreateJitSecretResourcePool detail param
type CreateJitSecretResourcePoolDetailParam struct {
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

// CreateJitSecretResourcePoolParam CreateJitSecretResourcePool request param
type CreateJitSecretResourcePoolParam struct {
	BaseParam
	Params CreateJitSecretResourcePoolDetailParam `json:"params"`
}
