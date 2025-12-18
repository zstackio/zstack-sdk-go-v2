// Copyright (c) ZStack.io, Inc.

package param

// AddAppBuildSystemDetailParam AddAppBuildSystem详细参数
type AddAppBuildSystemDetailParam struct {
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"storageType,omitempty"`
	rest string `json:"username" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"hostname" validate:"required"` // 必填
	rest int `json:"sshPort,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAppBuildSystemParam AddAppBuildSystem请求参数
type AddAppBuildSystemParam struct {
	BaseParam
	Params AddAppBuildSystemDetailParam `json:"params"` // 详细参数
}

