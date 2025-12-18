// Copyright (c) ZStack.io, Inc.

package param

// CreateCSPSecretResourcePoolDetailParam CreateCSPSecretResourcePool detail param
type CreateCSPSecretResourcePoolDetailParam struct {
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
	Params CreateCSPSecretResourcePoolDetailParam `json:"params"`
}
