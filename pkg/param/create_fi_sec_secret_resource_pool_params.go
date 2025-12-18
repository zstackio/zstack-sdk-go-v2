// Copyright (c) ZStack.io, Inc.

package param

// CreateFiSecSecretResourcePoolDetailParam CreateFiSecSecretResourcePool detail param
type CreateFiSecSecretResourcePoolDetailParam struct {
	KeyNum string `json:"keyNum,omitempty"`
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

// CreateFiSecSecretResourcePoolParam CreateFiSecSecretResourcePool request param
type CreateFiSecSecretResourcePoolParam struct {
	BaseParam
	Params CreateFiSecSecretResourcePoolDetailParam `json:"params"`
}
