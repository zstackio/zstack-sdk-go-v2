// Copyright (c) ZStack.io, Inc.

package param

// CreatePortMirrorSessionDetailParam CreatePortMirrorSession detail param
type CreatePortMirrorSessionDetailParam struct {
	PortMirrorUuid string `json:"portMirrorUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	SrcEndPoint string `json:"srcEndPoint" validate:"required"`
	DstEndPoint string `json:"dstEndPoint" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePortMirrorSessionParam CreatePortMirrorSession request param
type CreatePortMirrorSessionParam struct {
	BaseParam
	Params CreatePortMirrorSessionDetailParam `json:"params"`
}
