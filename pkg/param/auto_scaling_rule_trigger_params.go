// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteAutoScalingRuleTriggerParamDetail DeleteAutoScalingRuleTrigger detail param
type DeleteAutoScalingRuleTriggerParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingRuleTriggerParam DeleteAutoScalingRuleTrigger request param
type DeleteAutoScalingRuleTriggerParam struct {
	BaseParam
	DeleteAutoScalingRuleTrigger DeleteAutoScalingRuleTriggerParamDetail `json:"deleteAutoScalingRuleTrigger"`
}
