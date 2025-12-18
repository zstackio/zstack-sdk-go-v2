// Copyright (c) ZStack.io, Inc.

package param

// UpdatePrimaryStorageDetailParam UpdatePrimaryStorage detail param
type UpdatePrimaryStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
}

// UpdatePrimaryStorageParam UpdatePrimaryStorage request param
type UpdatePrimaryStorageParam struct {
	BaseParam
	Params UpdatePrimaryStorageDetailParam `json:"params"`
}
