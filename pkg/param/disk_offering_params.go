// Copyright (c) ZStack.io, Inc.

package param

// CreateDiskOfferingDetailParam CreateDiskOffering详细参数
type CreateDiskOfferingDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest int64 `json:"diskSize" validate:"required"` // 必填
	rest int `json:"sortKey,omitempty"`
	rest string `json:"allocationStrategy,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateDiskOfferingParam CreateDiskOffering请求参数
type CreateDiskOfferingParam struct {
	BaseParam
	Params CreateDiskOfferingDetailParam `json:"params"` // 详细参数
}

