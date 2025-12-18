// Copyright (c) ZStack.io, Inc.

package param

// GetSupportedIdentityModelsDetailParam GetSupportedIdentityModels detail param
type GetSupportedIdentityModelsDetailParam struct {
}

// GetSupportedIdentityModelsParam GetSupportedIdentityModels request param
type GetSupportedIdentityModelsParam struct {
	BaseParam
	Params GetSupportedIdentityModelsDetailParam `json:"params"`
}
