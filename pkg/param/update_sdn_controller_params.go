// Copyright (c) ZStack.io, Inc.

package param

// UpdateSdnControllerDetailParam UpdateSdnController detail param
type UpdateSdnControllerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateSdnControllerParam UpdateSdnController request param
type UpdateSdnControllerParam struct {
	BaseParam
	Params UpdateSdnControllerDetailParam `json:"params"`
}
