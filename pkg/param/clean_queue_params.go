// Copyright (c) ZStack.io, Inc.

package param

// CleanQueueDetailParam CleanQueue详细参数
type CleanQueueDetailParam struct {
	rest string `json:"signatureName" validate:"required"` // 必填
	rest int `json:"taskIndex,omitempty"`
	rest bool `json:"isCleanUp,omitempty"`
	rest bool `json:"isRunningTask,omitempty"`
	rest string `json:"managementiUuid,omitempty"`
}

// CleanQueueParam CleanQueue请求参数
type CleanQueueParam struct {
	BaseParam
	Params CleanQueueDetailParam `json:"params"` // 详细参数
}

