// Copyright (c) ZStack.io, Inc.

package param

// GetOAuth2TokenDetailParam GetOAuth2Token详细参数
type GetOAuth2TokenDetailParam struct {
}

// GetOAuth2TokenParam GetOAuth2Token请求参数
type GetOAuth2TokenParam struct {
	BaseParam
	Params GetOAuth2TokenDetailParam `json:"params"` // 详细参数
}

