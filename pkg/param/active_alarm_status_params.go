// Copyright (c) ZStack.io, Inc.

package param

// GetActiveAlarmStatusDetailParam GetActiveAlarmStatus详细参数
type GetActiveAlarmStatusDetailParam struct {
	rest string `json:"accountUuid" validate:"required"` // 必填
}

// GetActiveAlarmStatusParam GetActiveAlarmStatus请求参数
type GetActiveAlarmStatusParam struct {
	BaseParam
	Params GetActiveAlarmStatusDetailParam `json:"params"` // 详细参数
}

