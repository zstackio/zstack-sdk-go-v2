// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2VirtualIDTypeDetailParam ChangeIAM2VirtualIDType详细参数
type ChangeIAM2VirtualIDTypeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
}

// ChangeIAM2VirtualIDTypeParam ChangeIAM2VirtualIDType请求参数
type ChangeIAM2VirtualIDTypeParam struct {
	BaseParam
	Params ChangeIAM2VirtualIDTypeDetailParam `json:"params"` // 详细参数
}

