// Copyright (c) ZStack.io, Inc.

package param

// GetSupportedIdentityModelsDetailParam GetSupportedIdentityModels详细参数
type GetSupportedIdentityModelsDetailParam struct {
}

// GetSupportedIdentityModelsParam GetSupportedIdentityModels请求参数
type GetSupportedIdentityModelsParam struct {
	BaseParam
	Params GetSupportedIdentityModelsDetailParam `json:"params"` // 详细参数
}

