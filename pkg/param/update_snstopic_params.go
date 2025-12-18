// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSTopicDetailParam UpdateSNSTopic detail param
type UpdateSNSTopicDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Locale string `json:"locale,omitempty"`
}

// UpdateSNSTopicParam UpdateSNSTopic request param
type UpdateSNSTopicParam struct {
	BaseParam
	Params UpdateSNSTopicDetailParam `json:"params"`
}
