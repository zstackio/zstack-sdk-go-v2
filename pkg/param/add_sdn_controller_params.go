// Copyright (c) ZStack.io, Inc.

package param

// AddSdnControllerDetailParam AddSdnController详细参数
type AddSdnControllerDetailParam struct {
	rest string `json:"vendorType" validate:"required"` // 必填
	rest string `json:"vendorVersion,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"ip" validate:"required"` // 必填
	rest string `json:"userName,omitempty"`
	rest string `json:"password,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddSdnControllerParam AddSdnController请求参数
type AddSdnControllerParam struct {
	BaseParam
	Params AddSdnControllerDetailParam `json:"params"` // 详细参数
}

