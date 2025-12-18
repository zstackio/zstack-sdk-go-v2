// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2VirtualIDDetailParam CreateIAM2VirtualID详细参数
type CreateIAM2VirtualIDDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest []interface{} `json:"attributes,omitempty"`
	rest string `json:"projectUuid,omitempty"`
	rest string `json:"organizationUuid,omitempty"`
	rest bool `json:"withoutDefaultRole,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateIAM2VirtualIDParam CreateIAM2VirtualID请求参数
type CreateIAM2VirtualIDParam struct {
	BaseParam
	Params CreateIAM2VirtualIDDetailParam `json:"params"` // 详细参数
}

