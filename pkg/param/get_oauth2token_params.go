// Copyright (c) ZStack.io, Inc.

package param

// GetOAuth2TokenDetailParam GetOAuth2Token detail param
type GetOAuth2TokenDetailParam struct {
}

// GetOAuth2TokenParam GetOAuth2Token request param
type GetOAuth2TokenParam struct {
	BaseParam
	Params GetOAuth2TokenDetailParam `json:"params"`
}
