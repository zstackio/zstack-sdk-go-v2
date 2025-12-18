// Copyright (c) ZStack.io, Inc.

package param

// RefreshSSOServerTokenDetailParam RefreshSSOServerToken详细参数
type RefreshSSOServerTokenDetailParam struct {
	rest string `json:"token" validate:"required"` // 必填
	rest int64 `json:"duration,omitempty"`
}

// RefreshSSOServerTokenParam RefreshSSOServerToken请求参数
type RefreshSSOServerTokenParam struct {
	BaseParam
	Params RefreshSSOServerTokenDetailParam `json:"params"` // 详细参数
}

