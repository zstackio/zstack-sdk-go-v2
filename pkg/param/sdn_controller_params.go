// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// RemoveSdnControllerParamDetail RemoveSdnController detail param
type RemoveSdnControllerParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveSdnControllerParam RemoveSdnController request param
type RemoveSdnControllerParam struct {
	BaseParam
	Params RemoveSdnControllerParamDetail `json:"removeSdnController"`
}
// AddSdnControllerParamDetail AddSdnController detail param
type AddSdnControllerParamDetail struct {
	VendorType string `json:"vendorType" validate:"required"`
	VendorVersion *string `json:"vendorVersion,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Ip string `json:"ip" validate:"required"`
	UserName *string `json:"userName,omitempty"`
	Password *string `json:"password,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSdnControllerParam AddSdnController request param
type AddSdnControllerParam struct {
	BaseParam
	Params AddSdnControllerParamDetail `json:"params"`
}
// UpdateSdnControllerParamDetail UpdateSdnController detail param
type UpdateSdnControllerParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateSdnControllerParam UpdateSdnController request param
type UpdateSdnControllerParam struct {
	BaseParam
	Params UpdateSdnControllerParamDetail `json:"updateSdnController"`
}
// ChangeSdnControllerParamDetail ChangeSdnController detail param
type ChangeSdnControllerParamDetail struct {
	UserName *string `json:"userName,omitempty"`
	Password *string `json:"password,omitempty"`
	VlanRanges []string `json:"vlanRanges,omitempty"`
}

// ChangeSdnControllerParam ChangeSdnController request param
type ChangeSdnControllerParam struct {
	BaseParam
	Params ChangeSdnControllerParamDetail `json:"changeSdnController"`
}
// ReconnectSdnControllerParamDetail ReconnectSdnController detail param
type ReconnectSdnControllerParamDetail struct {
}

// ReconnectSdnControllerParam ReconnectSdnController request param
type ReconnectSdnControllerParam struct {
	BaseParam
	Params ReconnectSdnControllerParamDetail `json:"reconnectSdnController"`
}
