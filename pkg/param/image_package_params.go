// Copyright (c) ZStack.io, Inc.

package param

// DeleteImagePackageDetailParam DeleteImagePackage详细参数
type DeleteImagePackageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteImagePackageParam DeleteImagePackage请求参数
type DeleteImagePackageParam struct {
	BaseParam
	Params DeleteImagePackageDetailParam `json:"params"` // 详细参数
}

