// Copyright (c) ZStack.io, Inc.

package param

// ReloadExternalServiceDetailParam ReloadExternalService detail param
type ReloadExternalServiceDetailParam struct {
	Name string `json:"name" validate:"required"`
}

// ReloadExternalServiceParam ReloadExternalService request param
type ReloadExternalServiceParam struct {
	BaseParam
	Params ReloadExternalServiceDetailParam `json:"params"`
}
