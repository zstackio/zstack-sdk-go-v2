// Copyright (c) ZStack.io, Inc.

package param

// RecoverImageDetailParam RecoverImage detail param
type RecoverImageDetailParam struct {
	ImageUuid string `json:"imageUuid" validate:"required"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
}

// RecoverImageParam RecoverImage request param
type RecoverImageParam struct {
	BaseParam
	Params RecoverImageDetailParam `json:"params"`
}
