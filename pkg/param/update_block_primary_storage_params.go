// Copyright (c) ZStack.io, Inc.

package param

// UpdateBlockPrimaryStorageDetailParam UpdateBlockPrimaryStorage detail param
type UpdateBlockPrimaryStorageDetailParam struct {
	VendorName string `json:"vendorName,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
}

// UpdateBlockPrimaryStorageParam UpdateBlockPrimaryStorage request param
type UpdateBlockPrimaryStorageParam struct {
	BaseParam
	Params UpdateBlockPrimaryStorageDetailParam `json:"params"`
}
