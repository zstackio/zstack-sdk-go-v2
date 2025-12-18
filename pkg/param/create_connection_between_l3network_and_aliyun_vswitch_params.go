// Copyright (c) ZStack.io, Inc.

package param

// CreateConnectionBetweenL3NetworkAndAliyunVSwitchDetailParam CreateConnectionBetweenL3NetworkAndAliyunVSwitch detail param
type CreateConnectionBetweenL3NetworkAndAliyunVSwitchDetailParam struct {
	L3networkUuid string `json:"l3networkUuid" validate:"required"`
	VpcUuid string `json:"vpcUuid" validate:"required"`
	VbrUuid string `json:"vbrUuid" validate:"required"`
	CpeIp string `json:"cpeIp" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Direction string `json:"direction" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateConnectionBetweenL3NetworkAndAliyunVSwitchParam CreateConnectionBetweenL3NetworkAndAliyunVSwitch request param
type CreateConnectionBetweenL3NetworkAndAliyunVSwitchParam struct {
	BaseParam
	Params CreateConnectionBetweenL3NetworkAndAliyunVSwitchDetailParam `json:"params"`
}
