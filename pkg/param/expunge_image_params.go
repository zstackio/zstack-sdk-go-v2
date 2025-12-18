// Copyright (c) ZStack.io, Inc.

package param

// ExpungeImageDetailParam ExpungeImage detail param
type ExpungeImageDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
}

// ExpungeImageParam ExpungeImage request param
type ExpungeImageParam struct {
	BaseParam
	Params ExpungeImageDetailParam `json:"params"`
}
