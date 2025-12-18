// Copyright (c) ZStack.io, Inc.

package param

// AddIntegrityResourceDetailParam AddIntegrityResource详细参数
type AddIntegrityResourceDetailParam struct {
	rest string `json:"resourceType" validate:"required"` // 必填
	rest int `json:"integrityResourceDataRangeInDays,omitempty"`
}

// AddIntegrityResourceParam AddIntegrityResource请求参数
type AddIntegrityResourceParam struct {
	BaseParam
	Params AddIntegrityResourceDetailParam `json:"params"` // 详细参数
}

