// Copyright (c) ZStack.io, Inc.

package param

// AddLogConfigurationDetailParam AddLogConfiguration detail param
type AddLogConfigurationDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	Level string `json:"level,omitempty"`
	Configuration string `json:"configuration" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddLogConfigurationParam AddLogConfiguration request param
type AddLogConfigurationParam struct {
	BaseParam
	Params AddLogConfigurationDetailParam `json:"params"`
}
