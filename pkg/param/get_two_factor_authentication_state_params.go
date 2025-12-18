// Copyright (c) ZStack.io, Inc.

package param

// GetTwoFactorAuthenticationStateDetailParam GetTwoFactorAuthenticationState detail param
type GetTwoFactorAuthenticationStateDetailParam struct {
}

// GetTwoFactorAuthenticationStateParam GetTwoFactorAuthenticationState request param
type GetTwoFactorAuthenticationStateParam struct {
	BaseParam
	Params GetTwoFactorAuthenticationStateDetailParam `json:"params"`
}
