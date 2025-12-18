// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunMountTargetDetailParam UpdateAliyunMountTarget详细参数
type UpdateAliyunMountTargetDetailParam struct {
	rest string `json:"accessGroupUuid" validate:"required"` // 必填
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateAliyunMountTargetParam UpdateAliyunMountTarget请求参数
type UpdateAliyunMountTargetParam struct {
	BaseParam
	Params UpdateAliyunMountTargetDetailParam `json:"params"` // 详细参数
}

