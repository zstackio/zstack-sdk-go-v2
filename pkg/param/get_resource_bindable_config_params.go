// Copyright (c) ZStack.io, Inc.

package param

// GetResourceBindableConfigDetailParam GetResourceBindableConfig detail param
type GetResourceBindableConfigDetailParam struct {
	Category string `json:"category,omitempty"`
}

// GetResourceBindableConfigParam GetResourceBindableConfig request param
type GetResourceBindableConfigParam struct {
	BaseParam
	Params GetResourceBindableConfigDetailParam `json:"params"`
}
