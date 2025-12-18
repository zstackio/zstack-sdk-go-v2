// Copyright (c) ZStack.io, Inc.

package param

// GetAliyunNasAccessGroupRemoteDetailParam GetAliyunNasAccessGroupRemote详细参数
type GetAliyunNasAccessGroupRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"groupName,omitempty"`
}

// GetAliyunNasAccessGroupRemoteParam GetAliyunNasAccessGroupRemote请求参数
type GetAliyunNasAccessGroupRemoteParam struct {
	BaseParam
	Params GetAliyunNasAccessGroupRemoteDetailParam `json:"params"` // 详细参数
}

