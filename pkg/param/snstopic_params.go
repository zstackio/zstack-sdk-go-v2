// Copyright (c) ZStack.io, Inc.

package param

// DeleteSNSTopicDetailParam DeleteSNSTopic详细参数
type DeleteSNSTopicDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteSNSTopicParam DeleteSNSTopic请求参数
type DeleteSNSTopicParam struct {
	BaseParam
	Params DeleteSNSTopicDetailParam `json:"params"` // 详细参数
}

