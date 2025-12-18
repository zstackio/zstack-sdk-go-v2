// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunMountTargetDetailParam UpdateAliyunMountTarget detail param
type UpdateAliyunMountTargetDetailParam struct {
	AccessGroupUuid string `json:"accessGroupUuid" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateAliyunMountTargetParam UpdateAliyunMountTarget request param
type UpdateAliyunMountTargetParam struct {
	BaseParam
	Params UpdateAliyunMountTargetDetailParam `json:"params"`
}
