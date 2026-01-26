// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateFiSecSecretResourcePoolParamDetail CreateFiSecSecretResourcePool detail param
type CreateFiSecSecretResourcePoolParamDetail struct {
	KeyNum *string `json:"keyNum,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Model *string `json:"model,omitempty"`
	Ability *string `json:"ability,omitempty"`
	Type string `json:"type" validate:"required"`
	HeartbeatInterval int `json:"heartbeatInterval" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFiSecSecretResourcePoolParam CreateFiSecSecretResourcePool request param
type CreateFiSecSecretResourcePoolParam struct {
	BaseParam
	Params CreateFiSecSecretResourcePoolParamDetail `json:"params"`
}
// UpdateFiSecSecretResourcePoolParamDetail UpdateFiSecSecretResourcePool detail param
type UpdateFiSecSecretResourcePoolParamDetail struct {
	KeyNum *string `json:"keyNum,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Model *string `json:"model,omitempty"`
	HeartbeatInterval *int `json:"heartbeatInterval,omitempty"`
}

// UpdateFiSecSecretResourcePoolParam UpdateFiSecSecretResourcePool request param
type UpdateFiSecSecretResourcePoolParam struct {
	BaseParam
	Params UpdateFiSecSecretResourcePoolParamDetail `json:"updateFiSecSecretResourcePool"`
}
