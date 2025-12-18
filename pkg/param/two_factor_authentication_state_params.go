// Copyright (c) ZStack.io, Inc.

package param

// GetTwoFactorAuthenticationStateDetailParam GetTwoFactorAuthenticationState详细参数
type GetTwoFactorAuthenticationStateDetailParam struct {
}

// GetTwoFactorAuthenticationStateParam GetTwoFactorAuthenticationState请求参数
type GetTwoFactorAuthenticationStateParam struct {
	BaseParam
	Params GetTwoFactorAuthenticationStateDetailParam `json:"params"` // 详细参数
}

