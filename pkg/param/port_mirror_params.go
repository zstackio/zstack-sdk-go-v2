// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreatePortMirrorParamDetail CreatePortMirror detail param
type CreatePortMirrorParamDetail struct {
	MirrorNetworkUuid string `json:"mirrorNetworkUuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	StateEvent string `json:"stateEvent,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePortMirrorParam CreatePortMirror request param
type CreatePortMirrorParam struct {
	BaseParam
	CreatePortMirror CreatePortMirrorParamDetail `json:"createPortMirror"`
}
// DeletePortMirrorParamDetail DeletePortMirror detail param
type DeletePortMirrorParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePortMirrorParam DeletePortMirror request param
type DeletePortMirrorParam struct {
	BaseParam
	DeletePortMirror DeletePortMirrorParamDetail `json:"deletePortMirror"`
}
// UpdatePortMirrorParamDetail UpdatePortMirror detail param
type UpdatePortMirrorParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdatePortMirrorParam UpdatePortMirror request param
type UpdatePortMirrorParam struct {
	BaseParam
	UpdatePortMirror UpdatePortMirrorParamDetail `json:"updatePortMirror"`
}
