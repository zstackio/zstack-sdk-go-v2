// Copyright (c) ZStack.io, Inc.

package param

// ApplyDRSAdviceDetailParam ApplyDRSAdvice详细参数
type ApplyDRSAdviceDetailParam struct {
	rest string `json:"adviceUuid" validate:"required"` // 必填
}

// ApplyDRSAdviceParam ApplyDRSAdvice请求参数
type ApplyDRSAdviceParam struct {
	BaseParam
	Params ApplyDRSAdviceDetailParam `json:"params"` // 详细参数
}

