// Copyright (c) ZStack.io, Inc.

package param

// StartEcsInstanceDetailParam StartEcsInstance详细参数
type StartEcsInstanceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// StartEcsInstanceParam StartEcsInstance请求参数
type StartEcsInstanceParam struct {
	BaseParam
	Params StartEcsInstanceDetailParam `json:"params"` // 详细参数
}

