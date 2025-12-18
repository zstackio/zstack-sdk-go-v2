// Copyright (c) ZStack.io, Inc.

package param

// BootstrapMiniHostDetailParam BootstrapMiniHost detail param
type BootstrapMiniHostDetailParam struct {
	Local interface{} `json:"local" validate:"required"`
	Peer interface{} `json:"peer" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// BootstrapMiniHostParam BootstrapMiniHost request param
type BootstrapMiniHostParam struct {
	BaseParam
	Params BootstrapMiniHostDetailParam `json:"params"`
}
