// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunEbsPrimaryStorageDetailParam AddAliyunEbsPrimaryStorage详细参数
type AddAliyunEbsPrimaryStorageDetailParam struct {
	rest string `json:"panguPartitionUuid,omitempty"`
	rest string `json:"identityZoneUuid,omitempty"`
	rest string `json:"defaultIoType,omitempty"`
	rest string `json:"tdcConfigContent" validate:"required"` // 必填
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAliyunEbsPrimaryStorageParam AddAliyunEbsPrimaryStorage请求参数
type AddAliyunEbsPrimaryStorageParam struct {
	BaseParam
	Params AddAliyunEbsPrimaryStorageDetailParam `json:"params"` // 详细参数
}

