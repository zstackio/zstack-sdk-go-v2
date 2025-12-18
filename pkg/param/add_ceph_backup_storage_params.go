// Copyright (c) ZStack.io, Inc.

package param

// AddCephBackupStorageDetailParam AddCephBackupStorage详细参数
type AddCephBackupStorageDetailParam struct {
	rest []string `json:"monUrls" validate:"required"` // 必填
	rest string `json:"poolName,omitempty"`
	rest string `json:"url,omitempty"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest bool `json:"importImages,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddCephBackupStorageParam AddCephBackupStorage请求参数
type AddCephBackupStorageParam struct {
	BaseParam
	Params AddCephBackupStorageDetailParam `json:"params"` // 详细参数
}

