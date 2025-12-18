// Copyright (c) ZStack.io, Inc.

package param

// AddDisasterImageStoreBackupStorageDetailParam AddDisasterImageStoreBackupStorage详细参数
type AddDisasterImageStoreBackupStorageDetailParam struct {
	rest string `json:"attachPoint,omitempty"`
	rest string `json:"endPoint,omitempty"`
	rest string `json:"hostname" validate:"required"` // 必填
	rest string `json:"username" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest int `json:"sshPort,omitempty"`
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest bool `json:"importImages,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddDisasterImageStoreBackupStorageParam AddDisasterImageStoreBackupStorage请求参数
type AddDisasterImageStoreBackupStorageParam struct {
	BaseParam
	Params AddDisasterImageStoreBackupStorageDetailParam `json:"params"` // 详细参数
}

