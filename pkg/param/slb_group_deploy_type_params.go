// Copyright (c) ZStack.io, Inc.

package param

// ChangeSlbGroupDeployTypeDetailParam ChangeSlbGroupDeployType详细参数
type ChangeSlbGroupDeployTypeDetailParam struct {
	rest string `json:"slbGroupUuid" validate:"required"` // 必填
	rest string `json:"deployType" validate:"required"` // 必填
}

// ChangeSlbGroupDeployTypeParam ChangeSlbGroupDeployType请求参数
type ChangeSlbGroupDeployTypeParam struct {
	BaseParam
	Params ChangeSlbGroupDeployTypeDetailParam `json:"params"` // 详细参数
}

