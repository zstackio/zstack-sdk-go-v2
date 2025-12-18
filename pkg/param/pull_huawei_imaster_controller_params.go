// Copyright (c) ZStack.io, Inc.

package param

// PullHuaweiIMasterControllerDetailParam PullHuaweiIMasterController详细参数
type PullHuaweiIMasterControllerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"pullSwitch,omitempty"`
}

// PullHuaweiIMasterControllerParam PullHuaweiIMasterController请求参数
type PullHuaweiIMasterControllerParam struct {
	BaseParam
	Params PullHuaweiIMasterControllerDetailParam `json:"params"` // 详细参数
}

