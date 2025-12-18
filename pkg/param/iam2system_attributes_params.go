// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2SystemAttributesDetailParam GetIAM2SystemAttributes详细参数
type GetIAM2SystemAttributesDetailParam struct {
}

// GetIAM2SystemAttributesParam GetIAM2SystemAttributes请求参数
type GetIAM2SystemAttributesParam struct {
	BaseParam
	Params GetIAM2SystemAttributesDetailParam `json:"params"` // 详细参数
}

