// Copyright (c) ZStack.io, Inc.

package param

// UpdateAlertDataAckDetailParam UpdateAlertDataAck detail param
type UpdateAlertDataAckDetailParam struct {
	AlertDataUuid string `json:"alertDataUuid" validate:"required"`
	ResumeAlert bool `json:"resumeAlert,omitempty"`
}

// UpdateAlertDataAckParam UpdateAlertDataAck request param
type UpdateAlertDataAckParam struct {
	BaseParam
	Params UpdateAlertDataAckDetailParam `json:"params"`
}
