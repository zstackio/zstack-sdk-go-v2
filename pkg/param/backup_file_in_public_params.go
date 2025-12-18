// Copyright (c) ZStack.io, Inc.

package param

// DeleteBackupFileInPublicDetailParam DeleteBackupFileInPublic详细参数
type DeleteBackupFileInPublicDetailParam struct {
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"regionId" validate:"required"` // 必填
	rest string `json:"file" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteBackupFileInPublicParam DeleteBackupFileInPublic请求参数
type DeleteBackupFileInPublicParam struct {
	BaseParam
	Params DeleteBackupFileInPublicDetailParam `json:"params"` // 详细参数
}

