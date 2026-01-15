// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateCSPSecretResourcePoolParamDetail UpdateCSPSecretResourcePool detail param
type UpdateCSPSecretResourcePoolParamDetail struct {
	AppId string `json:"appId,omitempty"`
	AppKey string `json:"appKey,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	KeyId string `json:"keyId,omitempty"`
	UserId string `json:"userId,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdateCSPSecretResourcePoolParam UpdateCSPSecretResourcePool request param
type UpdateCSPSecretResourcePoolParam struct {
	BaseParam
	Params UpdateCSPSecretResourcePoolParamDetail `json:"updateCSPSecretResourcePool"`
}
// CreateCSPSecretResourcePoolParamDetail CreateCSPSecretResourcePool detail param
type CreateCSPSecretResourcePoolParamDetail struct {
	ManagementIp string `json:"managementIp" validate:"required"`
	Port int `json:"port" validate:"required"`
	AppId string `json:"appId" validate:"required"`
	AppKey string `json:"appKey" validate:"required"`
	KeyId string `json:"keyId" validate:"required"`
	Protocol string `json:"protocol,omitempty"`
	UserId string `json:"userId,omitempty"`
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

// CreateCSPSecretResourcePoolParam CreateCSPSecretResourcePool request param
type CreateCSPSecretResourcePoolParam struct {
	BaseParam
	Params CreateCSPSecretResourcePoolParamDetail `json:"createCSPSecretResourcePool"`
}
