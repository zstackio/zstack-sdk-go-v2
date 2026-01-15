// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateAlertDataAckParamDetail UpdateAlertDataAck detail param
type UpdateAlertDataAckParamDetail struct {
	AlertDataUuid string `json:"alertDataUuid" validate:"required"`
	ResumeAlert bool `json:"resumeAlert,omitempty"`
}

// UpdateAlertDataAckParam UpdateAlertDataAck request param
type UpdateAlertDataAckParam struct {
	BaseParam
	UpdateAlertDataAck UpdateAlertDataAckParamDetail `json:"updateAlertDataAck"`
}
