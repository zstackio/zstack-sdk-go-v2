// Copyright (c) ZStack.io, Inc.

package param

// CleanV2VConversionCacheDetailParam CleanV2VConversionCache详细参数
type CleanV2VConversionCacheDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// CleanV2VConversionCacheParam CleanV2VConversionCache请求参数
type CleanV2VConversionCacheParam struct {
	BaseParam
	Params CleanV2VConversionCacheDetailParam `json:"params"` // 详细参数
}

