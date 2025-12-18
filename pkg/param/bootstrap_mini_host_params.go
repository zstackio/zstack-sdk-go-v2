// Copyright (c) ZStack.io, Inc.

package param

// BootstrapMiniHostDetailParam BootstrapMiniHost详细参数
type BootstrapMiniHostDetailParam struct {
	rest interface{} `json:"local" validate:"required"` // 必填
	rest interface{} `json:"peer" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// BootstrapMiniHostParam BootstrapMiniHost请求参数
type BootstrapMiniHostParam struct {
	BaseParam
	Params BootstrapMiniHostDetailParam `json:"params"` // 详细参数
}

