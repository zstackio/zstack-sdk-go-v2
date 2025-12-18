// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunEbsPrimaryStorageDetailParam UpdateAliyunEbsPrimaryStorage详细参数
type UpdateAliyunEbsPrimaryStorageDetailParam struct {
	rest string `json:"panguAppName,omitempty"`
	rest string `json:"panguPartitionName,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"url,omitempty"`
}

// UpdateAliyunEbsPrimaryStorageParam UpdateAliyunEbsPrimaryStorage请求参数
type UpdateAliyunEbsPrimaryStorageParam struct {
	BaseParam
	Params UpdateAliyunEbsPrimaryStorageDetailParam `json:"params"` // 详细参数
}

