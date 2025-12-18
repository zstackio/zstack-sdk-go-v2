// Copyright (c) ZStack.io, Inc.

package param

// RefreshSSOServerTokenDetailParam RefreshSSOServerToken detail param
type RefreshSSOServerTokenDetailParam struct {
	Token string `json:"token" validate:"required"`
	Duration int64 `json:"duration,omitempty"`
}

// RefreshSSOServerTokenParam RefreshSSOServerToken request param
type RefreshSSOServerTokenParam struct {
	BaseParam
	Params RefreshSSOServerTokenDetailParam `json:"params"`
}
