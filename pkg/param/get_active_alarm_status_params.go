// Copyright (c) ZStack.io, Inc.

package param

// GetActiveAlarmStatusDetailParam GetActiveAlarmStatus detail param
type GetActiveAlarmStatusDetailParam struct {
	AccountUuid string `json:"accountUuid" validate:"required"`
}

// GetActiveAlarmStatusParam GetActiveAlarmStatus request param
type GetActiveAlarmStatusParam struct {
	BaseParam
	Params GetActiveAlarmStatusDetailParam `json:"params"`
}
