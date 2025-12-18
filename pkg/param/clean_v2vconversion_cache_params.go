// Copyright (c) ZStack.io, Inc.

package param

// CleanV2VConversionCacheDetailParam CleanV2VConversionCache detail param
type CleanV2VConversionCacheDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// CleanV2VConversionCacheParam CleanV2VConversionCache request param
type CleanV2VConversionCacheParam struct {
	BaseParam
	Params CleanV2VConversionCacheDetailParam `json:"params"`
}
