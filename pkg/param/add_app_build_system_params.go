// Copyright (c) ZStack.io, Inc.

package param

// AddAppBuildSystemDetailParam AddAppBuildSystem detail param
type AddAppBuildSystemDetailParam struct {
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	StorageType string `json:"storageType,omitempty"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Hostname string `json:"hostname" validate:"required"`
	SshPort int `json:"sshPort,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAppBuildSystemParam AddAppBuildSystem request param
type AddAppBuildSystemParam struct {
	BaseParam
	Params AddAppBuildSystemDetailParam `json:"params"`
}
