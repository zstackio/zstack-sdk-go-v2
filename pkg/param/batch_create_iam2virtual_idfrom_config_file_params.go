// Copyright (c) ZStack.io, Inc.

package param

// BatchCreateIAM2VirtualIDFromConfigFileDetailParam BatchCreateIAM2VirtualIDFromConfigFile详细参数
type BatchCreateIAM2VirtualIDFromConfigFileDetailParam struct {
	rest string `json:"virtualIDInfos" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// BatchCreateIAM2VirtualIDFromConfigFileParam BatchCreateIAM2VirtualIDFromConfigFile请求参数
type BatchCreateIAM2VirtualIDFromConfigFileParam struct {
	BaseParam
	Params BatchCreateIAM2VirtualIDFromConfigFileDetailParam `json:"params"` // 详细参数
}

