// Copyright (c) ZStack.io, Inc.

package param

// DeleteImageReplicationGroupDetailParam DeleteImageReplicationGroup detail param
type DeleteImageReplicationGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteImageReplicationGroupParam DeleteImageReplicationGroup request param
type DeleteImageReplicationGroupParam struct {
	BaseParam
	Params DeleteImageReplicationGroupDetailParam `json:"params"`
}
