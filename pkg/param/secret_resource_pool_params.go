// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteSecretResourcePoolParamDetail DeleteSecretResourcePool detail param
type DeleteSecretResourcePoolParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSecretResourcePoolParam DeleteSecretResourcePool request param
type DeleteSecretResourcePoolParam struct {
	BaseParam
	DeleteSecretResourcePool DeleteSecretResourcePoolParamDetail `json:"deleteSecretResourcePool"`
}
// UpdateSecretResourcePoolParamDetail UpdateSecretResourcePool detail param
type UpdateSecretResourcePoolParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdateSecretResourcePoolParam UpdateSecretResourcePool request param
type UpdateSecretResourcePoolParam struct {
	BaseParam
	UpdateSecretResourcePool UpdateSecretResourcePoolParamDetail `json:"updateSecretResourcePool"`
}
