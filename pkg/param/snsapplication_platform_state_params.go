// Copyright (c) ZStack.io, Inc.

package param

// ChangeSNSApplicationPlatformStateDetailParam ChangeSNSApplicationPlatformState详细参数
type ChangeSNSApplicationPlatformStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeSNSApplicationPlatformStateParam ChangeSNSApplicationPlatformState请求参数
type ChangeSNSApplicationPlatformStateParam struct {
	BaseParam
	Params ChangeSNSApplicationPlatformStateDetailParam `json:"params"` // 详细参数
}

