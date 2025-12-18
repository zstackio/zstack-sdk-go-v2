// Copyright (c) ZStack.io, Inc.

package param

// AddExternalBackupStorageDetailParam AddExternalBackupStorage详细参数
type AddExternalBackupStorageDetailParam struct {
	rest string `json:"identity" validate:"required"` // 必填
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest bool `json:"importImages,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddExternalBackupStorageParam AddExternalBackupStorage请求参数
type AddExternalBackupStorageParam struct {
	BaseParam
	Params AddExternalBackupStorageDetailParam `json:"params"` // 详细参数
}

