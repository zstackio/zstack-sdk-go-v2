// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageLicenseInfoDetailParam GetPrimaryStorageLicenseInfo detail param
type GetPrimaryStorageLicenseInfoDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetPrimaryStorageLicenseInfoParam GetPrimaryStorageLicenseInfo request param
type GetPrimaryStorageLicenseInfoParam struct {
	BaseParam
	Params GetPrimaryStorageLicenseInfoDetailParam `json:"params"`
}
