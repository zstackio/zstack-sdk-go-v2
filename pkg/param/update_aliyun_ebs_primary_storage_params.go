// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunEbsPrimaryStorageDetailParam UpdateAliyunEbsPrimaryStorage detail param
type UpdateAliyunEbsPrimaryStorageDetailParam struct {
	PanguAppName string `json:"panguAppName,omitempty"`
	PanguPartitionName string `json:"panguPartitionName,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
}

// UpdateAliyunEbsPrimaryStorageParam UpdateAliyunEbsPrimaryStorage request param
type UpdateAliyunEbsPrimaryStorageParam struct {
	BaseParam
	Params UpdateAliyunEbsPrimaryStorageDetailParam `json:"params"`
}
