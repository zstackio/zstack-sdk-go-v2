// Copyright (c) ZStack.io, Inc.

package param

// UpdateExternalPrimaryStorageDetailParam UpdateExternalPrimaryStorage detail param
type UpdateExternalPrimaryStorageDetailParam struct {
	Config string `json:"config,omitempty"`
	DefaultProtocol string `json:"defaultProtocol,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
}

// UpdateExternalPrimaryStorageParam UpdateExternalPrimaryStorage request param
type UpdateExternalPrimaryStorageParam struct {
	BaseParam
	Params UpdateExternalPrimaryStorageDetailParam `json:"params"`
}
