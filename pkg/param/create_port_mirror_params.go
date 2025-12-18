// Copyright (c) ZStack.io, Inc.

package param

// CreatePortMirrorDetailParam CreatePortMirror detail param
type CreatePortMirrorDetailParam struct {
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
	Params CreatePortMirrorDetailParam `json:"params"`
}
