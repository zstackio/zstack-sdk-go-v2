// Copyright (c) ZStack.io, Inc.

package param

// CreateVRouterOspfAreaDetailParam CreateVRouterOspfArea详细参数
type CreateVRouterOspfAreaDetailParam struct {
	rest string `json:"areaId" validate:"required"` // 必填
	rest string `json:"areaAuth,omitempty"`
	rest string `json:"areaType,omitempty"`
	rest string `json:"password,omitempty"`
	rest int `json:"keyId,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVRouterOspfAreaParam CreateVRouterOspfArea请求参数
type CreateVRouterOspfAreaParam struct {
	BaseParam
	Params CreateVRouterOspfAreaDetailParam `json:"params"` // 详细参数
}

