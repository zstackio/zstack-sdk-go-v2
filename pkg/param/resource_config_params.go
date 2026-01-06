// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// GetResourceConfigParamDetail GetResourceConfig detail param
type GetResourceConfigParamDetail struct {
	Category string `json:"category" validate:"required"`
	Name string `json:"name" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
}

// GetResourceConfigParam GetResourceConfig request param
type GetResourceConfigParam struct {
	BaseParam
	Params GetResourceConfigParamDetail `json:"params"`
}
// UpdateResourceConfigParamDetail UpdateResourceConfig detail param
type UpdateResourceConfigParamDetail struct {
	Category string `json:"category" validate:"required"`
	Name string `json:"name" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateResourceConfigParam UpdateResourceConfig request param
type UpdateResourceConfigParam struct {
	BaseParam
	Params UpdateResourceConfigParamDetail `json:"params"`
}
// DeleteResourceConfigParamDetail DeleteResourceConfig detail param
type DeleteResourceConfigParamDetail struct {
	Category string `json:"category" validate:"required"`
	Name string `json:"name" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteResourceConfigParam DeleteResourceConfig request param
type DeleteResourceConfigParam struct {
	BaseParam
	Params DeleteResourceConfigParamDetail `json:"params"`
}
