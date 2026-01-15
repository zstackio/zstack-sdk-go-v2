// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateSanSecSecretResourcePoolParamDetail CreateSanSecSecretResourcePool detail param
type CreateSanSecSecretResourcePoolParamDetail struct {
	KeyIndex int `json:"keyIndex,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Sm3Key string `json:"sm3Key,omitempty"`
	Sm4Key string `json:"sm4Key,omitempty"`
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

// CreateSanSecSecretResourcePoolParam CreateSanSecSecretResourcePool request param
type CreateSanSecSecretResourcePoolParam struct {
	BaseParam
	CreateSanSecSecretResourcePool CreateSanSecSecretResourcePoolParamDetail `json:"createSanSecSecretResourcePool"`
}
// UpdateSanSecSecretResourcePoolParamDetail UpdateSanSecSecretResourcePool detail param
type UpdateSanSecSecretResourcePoolParamDetail struct {
	KeyIndex int `json:"keyIndex,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Sm3Key string `json:"sm3Key,omitempty"`
	Sm4Key string `json:"sm4Key,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdateSanSecSecretResourcePoolParam UpdateSanSecSecretResourcePool request param
type UpdateSanSecSecretResourcePoolParam struct {
	BaseParam
	UpdateSanSecSecretResourcePool UpdateSanSecSecretResourcePoolParamDetail `json:"updateSanSecSecretResourcePool"`
}
