// Copyright (c) ZStack.io, Inc.

package param

// DeleteSNSTopicDetailParam DeleteSNSTopic detail param
type DeleteSNSTopicDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSNSTopicParam DeleteSNSTopic request param
type DeleteSNSTopicParam struct {
	BaseParam
	Params DeleteSNSTopicDetailParam `json:"params"`
}
