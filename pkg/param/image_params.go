// Copyright (c) ZStack.io, Inc.

package param

// ExpungeImageDetailParam ExpungeImage详细参数
type ExpungeImageDetailParam struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest []string `json:"backupStorageUuids,omitempty"`
}

// ExpungeImageParam ExpungeImage请求参数
type ExpungeImageParam struct {
	BaseParam
	Params ExpungeImageDetailParam `json:"params"` // 详细参数
}

