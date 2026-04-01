// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreatePluginSecretResourcePoolParamDetail CreatePluginSecretResourcePool detail param
type CreatePluginSecretResourcePoolParamDetail struct {
	Properties map[string]string `json:"properties,omitempty"`
	PluginDriverUuid *string `json:"pluginDriverUuid,omitempty"`
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

// CreatePluginSecretResourcePoolParam CreatePluginSecretResourcePool request param
type CreatePluginSecretResourcePoolParam struct {
	BaseParam
	Params CreatePluginSecretResourcePoolParamDetail `json:"params"`
}
// UpdatePluginSecretResourcePoolParamDetail UpdatePluginSecretResourcePool detail param
type UpdatePluginSecretResourcePoolParamDetail struct {
	Properties map[string]string `json:"properties,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Model *string `json:"model,omitempty"`
	HeartbeatInterval *int `json:"heartbeatInterval,omitempty"`
}

// UpdatePluginSecretResourcePoolParam UpdatePluginSecretResourcePool request param
type UpdatePluginSecretResourcePoolParam struct {
	BaseParam
	Params UpdatePluginSecretResourcePoolParamDetail `json:"updatePluginSecretResourcePool"`
}
