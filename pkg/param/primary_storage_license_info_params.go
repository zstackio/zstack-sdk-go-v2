// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageLicenseInfoDetailParam GetPrimaryStorageLicenseInfo详细参数
type GetPrimaryStorageLicenseInfoDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetPrimaryStorageLicenseInfoParam GetPrimaryStorageLicenseInfo请求参数
type GetPrimaryStorageLicenseInfoParam struct {
	BaseParam
	Params GetPrimaryStorageLicenseInfoDetailParam `json:"params"` // 详细参数
}

