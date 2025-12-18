// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunNasAccessGroupDetailParam UpdateAliyunNasAccessGroup详细参数
type UpdateAliyunNasAccessGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
}

// UpdateAliyunNasAccessGroupParam UpdateAliyunNasAccessGroup请求参数
type UpdateAliyunNasAccessGroupParam struct {
	BaseParam
	Params UpdateAliyunNasAccessGroupDetailParam `json:"params"` // 详细参数
}

