// Copyright (c) ZStack.io, Inc.

package param

// ApplyDRSAdviceDetailParam ApplyDRSAdvice detail param
type ApplyDRSAdviceDetailParam struct {
	AdviceUuid string `json:"adviceUuid" validate:"required"`
}

// ApplyDRSAdviceParam ApplyDRSAdvice request param
type ApplyDRSAdviceParam struct {
	BaseParam
	Params ApplyDRSAdviceDetailParam `json:"params"`
}
