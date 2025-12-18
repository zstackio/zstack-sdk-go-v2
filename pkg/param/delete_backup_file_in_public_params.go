// Copyright (c) ZStack.io, Inc.

package param

// DeleteBackupFileInPublicDetailParam DeleteBackupFileInPublic detail param
type DeleteBackupFileInPublicDetailParam struct {
	Type string `json:"type" validate:"required"`
	RegionId string `json:"regionId" validate:"required"`
	File string `json:"file" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBackupFileInPublicParam DeleteBackupFileInPublic request param
type DeleteBackupFileInPublicParam struct {
	BaseParam
	Params DeleteBackupFileInPublicDetailParam `json:"params"`
}
