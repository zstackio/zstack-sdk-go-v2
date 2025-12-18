// Copyright (c) ZStack.io, Inc.

package param

// CreateImageGroupFromVmInstanceDetailParam CreateImageGroupFromVmInstance详细参数
type CreateImageGroupFromVmInstanceDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateImageGroupFromVmInstanceParam CreateImageGroupFromVmInstance请求参数
type CreateImageGroupFromVmInstanceParam struct {
	BaseParam
	Params CreateImageGroupFromVmInstanceDetailParam `json:"params"` // 详细参数
}

