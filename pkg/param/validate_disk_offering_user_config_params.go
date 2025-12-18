// Copyright (c) ZStack.io, Inc.

package param

// ValidateDiskOfferingUserConfigDetailParam ValidateDiskOfferingUserConfig detail param
type ValidateDiskOfferingUserConfigDetailParam struct {
	Config string `json:"config" validate:"required"`
}

// ValidateDiskOfferingUserConfigParam ValidateDiskOfferingUserConfig request param
type ValidateDiskOfferingUserConfigParam struct {
	BaseParam
	Params ValidateDiskOfferingUserConfigDetailParam `json:"params"`
}
