// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2TickFlowCollectionDetailParam CreateIAM2TickFlowCollection详细参数
type CreateIAM2TickFlowCollectionDetailParam struct {
	rest []interface{} `json:"flows,omitempty"`
	rest string `json:"projectUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest bool `json:"isDefault,omitempty"`
	rest []string `json:"ticketTypeUuids,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateIAM2TickFlowCollectionParam CreateIAM2TickFlowCollection请求参数
type CreateIAM2TickFlowCollectionParam struct {
	BaseParam
	Params CreateIAM2TickFlowCollectionDetailParam `json:"params"` // 详细参数
}

