// Copyright (c) ZStack.io, Inc.

package param

// DeleteImagePackageDetailParam DeleteImagePackage detail param
type DeleteImagePackageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteImagePackageParam DeleteImagePackage request param
type DeleteImagePackageParam struct {
	BaseParam
	Params DeleteImagePackageDetailParam `json:"params"`
}
