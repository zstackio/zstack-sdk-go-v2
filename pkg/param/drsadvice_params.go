// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// ApplyDRSAdviceParamDetail ApplyDRSAdvice detail param
type ApplyDRSAdviceParamDetail struct {
	AdviceUuid string `json:"adviceUuid" validate:"required"`
}

// ApplyDRSAdviceParam ApplyDRSAdvice request param
type ApplyDRSAdviceParam struct {
	BaseParam
	ApplyDRSAdvice ApplyDRSAdviceParamDetail `json:"applyDRSAdvice"`
}
