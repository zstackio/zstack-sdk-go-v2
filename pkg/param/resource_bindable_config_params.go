// Copyright (c) ZStack.io, Inc.

package param

// GetResourceBindableConfigDetailParam GetResourceBindableConfig详细参数
type GetResourceBindableConfigDetailParam struct {
	rest string `json:"category,omitempty"`
}

// GetResourceBindableConfigParam GetResourceBindableConfig请求参数
type GetResourceBindableConfigParam struct {
	BaseParam
	Params GetResourceBindableConfigDetailParam `json:"params"` // 详细参数
}

