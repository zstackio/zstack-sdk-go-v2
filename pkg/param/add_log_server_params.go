// Copyright (c) ZStack.io, Inc.

package param

// AddLogServerDetailParam AddLogServer detail param
type AddLogServerDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Category string `json:"category" validate:"required"`
	Type string `json:"type" validate:"required"`
	Level string `json:"level,omitempty"`
	Configuration string `json:"configuration" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddLogServerParam AddLogServer request param
type AddLogServerParam struct {
	BaseParam
	Params AddLogServerDetailParam `json:"params"`
}
