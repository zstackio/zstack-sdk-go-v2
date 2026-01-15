// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateKoAlSecretResourcePoolParamDetail UpdateKoAlSecretResourcePool detail param
type UpdateKoAlSecretResourcePoolParamDetail struct {
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdateKoAlSecretResourcePoolParam UpdateKoAlSecretResourcePool request param
type UpdateKoAlSecretResourcePoolParam struct {
	BaseParam
	UpdateKoAlSecretResourcePool UpdateKoAlSecretResourcePoolParamDetail `json:"updateKoAlSecretResourcePool"`
}
// CreateKoAlSecretResourcePoolParamDetail CreateKoAlSecretResourcePool detail param
type CreateKoAlSecretResourcePoolParamDetail struct {
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`
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

// CreateKoAlSecretResourcePoolParam CreateKoAlSecretResourcePool request param
type CreateKoAlSecretResourcePoolParam struct {
	BaseParam
	CreateKoAlSecretResourcePool CreateKoAlSecretResourcePoolParamDetail `json:"createKoAlSecretResourcePool"`
}
