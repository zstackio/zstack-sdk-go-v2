// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunNasPrimaryStorageDetailParam AddAliyunNasPrimaryStorage详细参数
type AddAliyunNasPrimaryStorageDetailParam struct {
	rest string `json:"nasUuid" validate:"required"` // 必填
	rest string `json:"accessGroupUuid" validate:"required"` // 必填
	rest string `json:"vSwitchUuid,omitempty"`
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasPrimaryStorageParam AddAliyunNasPrimaryStorage请求参数
type AddAliyunNasPrimaryStorageParam struct {
	BaseParam
	Params AddAliyunNasPrimaryStorageDetailParam `json:"params"` // 详细参数
}

