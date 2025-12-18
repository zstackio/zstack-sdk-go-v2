// Copyright (c) ZStack.io, Inc.

package param

// UpdatePublishAppDetailParam UpdatePublishApp detail param
type UpdatePublishAppDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdatePublishAppParam UpdatePublishApp request param
type UpdatePublishAppParam struct {
	BaseParam
	Params UpdatePublishAppDetailParam `json:"params"`
}
