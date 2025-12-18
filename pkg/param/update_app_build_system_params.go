// Copyright (c) ZStack.io, Inc.

package param

// UpdateAppBuildSystemDetailParam UpdateAppBuildSystem detail param
type UpdateAppBuildSystemDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
}

// UpdateAppBuildSystemParam UpdateAppBuildSystem request param
type UpdateAppBuildSystemParam struct {
	BaseParam
	Params UpdateAppBuildSystemDetailParam `json:"params"`
}
