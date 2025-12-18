// Copyright (c) ZStack.io, Inc.

package param

// UpdateThirdpartyPlatformDetailParam UpdateThirdpartyPlatform detail param
type UpdateThirdpartyPlatformDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Template string `json:"template,omitempty"`
	Url string `json:"url,omitempty"`
	StateEvent string `json:"stateEvent,omitempty"`
	LastSyncDateMills int64 `json:"lastSyncDateMills,omitempty"`
}

// UpdateThirdpartyPlatformParam UpdateThirdpartyPlatform request param
type UpdateThirdpartyPlatformParam struct {
	BaseParam
	Params UpdateThirdpartyPlatformDetailParam `json:"params"`
}
