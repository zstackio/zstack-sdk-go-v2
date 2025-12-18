// Copyright (c) ZStack.io, Inc.

package param

// UpdateImagePackageDetailParam UpdateImagePackage detail param
type UpdateImagePackageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateImagePackageParam UpdateImagePackage request param
type UpdateImagePackageParam struct {
	BaseParam
	Params UpdateImagePackageDetailParam `json:"params"`
}
