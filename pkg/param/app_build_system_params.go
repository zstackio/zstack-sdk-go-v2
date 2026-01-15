// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddAppBuildSystemParamDetail AddAppBuildSystem detail param
type AddAppBuildSystemParamDetail struct {
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
	Params AddAppBuildSystemParamDetail `json:"addAppBuildSystem"`
}
// ReconnectAppBuildSystemParamDetail ReconnectAppBuildSystem detail param
type ReconnectAppBuildSystemParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectAppBuildSystemParam ReconnectAppBuildSystem request param
type ReconnectAppBuildSystemParam struct {
	BaseParam
	Params ReconnectAppBuildSystemParamDetail `json:"reconnectAppBuildSystem"`
}
// UpdateAppBuildSystemParamDetail UpdateAppBuildSystem detail param
type UpdateAppBuildSystemParamDetail struct {
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
	Params UpdateAppBuildSystemParamDetail `json:"updateAppBuildSystem"`
}
// DeleteAppBuildSystemParamDetail DeleteAppBuildSystem detail param
type DeleteAppBuildSystemParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAppBuildSystemParam DeleteAppBuildSystem request param
type DeleteAppBuildSystemParam struct {
	BaseParam
	Params DeleteAppBuildSystemParamDetail `json:"deleteAppBuildSystem"`
}
