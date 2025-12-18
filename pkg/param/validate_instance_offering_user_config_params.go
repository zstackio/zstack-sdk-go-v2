// Copyright (c) ZStack.io, Inc.

package param

// ValidateInstanceOfferingUserConfigDetailParam ValidateInstanceOfferingUserConfig detail param
type ValidateInstanceOfferingUserConfigDetailParam struct {
	Config string `json:"config" validate:"required"`
}

// ValidateInstanceOfferingUserConfigParam ValidateInstanceOfferingUserConfig request param
type ValidateInstanceOfferingUserConfigParam struct {
	BaseParam
	Params ValidateInstanceOfferingUserConfigDetailParam `json:"params"`
}
