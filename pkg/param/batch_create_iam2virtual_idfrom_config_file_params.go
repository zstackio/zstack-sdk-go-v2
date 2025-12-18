// Copyright (c) ZStack.io, Inc.

package param

// BatchCreateIAM2VirtualIDFromConfigFileDetailParam BatchCreateIAM2VirtualIDFromConfigFile detail param
type BatchCreateIAM2VirtualIDFromConfigFileDetailParam struct {
	VirtualIDInfos string `json:"virtualIDInfos" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// BatchCreateIAM2VirtualIDFromConfigFileParam BatchCreateIAM2VirtualIDFromConfigFile request param
type BatchCreateIAM2VirtualIDFromConfigFileParam struct {
	BaseParam
	Params BatchCreateIAM2VirtualIDFromConfigFileDetailParam `json:"params"`
}
