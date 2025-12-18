// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelsDetailParam DeleteModels detail param
type DeleteModelsDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteModelsParam DeleteModels request param
type DeleteModelsParam struct {
	BaseParam
	Params DeleteModelsDetailParam `json:"params"`
}
