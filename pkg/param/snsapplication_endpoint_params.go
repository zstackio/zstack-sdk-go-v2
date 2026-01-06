// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateSNSApplicationEndpointParamDetail UpdateSNSApplicationEndpoint detail param
type UpdateSNSApplicationEndpointParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
}

// UpdateSNSApplicationEndpointParam UpdateSNSApplicationEndpoint request param
type UpdateSNSApplicationEndpointParam struct {
	BaseParam
	Params UpdateSNSApplicationEndpointParamDetail `json:"params"`
}
// DeleteSNSApplicationEndpointParamDetail DeleteSNSApplicationEndpoint detail param
type DeleteSNSApplicationEndpointParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSNSApplicationEndpointParam DeleteSNSApplicationEndpoint request param
type DeleteSNSApplicationEndpointParam struct {
	BaseParam
	Params DeleteSNSApplicationEndpointParamDetail `json:"params"`
}
