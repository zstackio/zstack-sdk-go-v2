// Copyright (c) ZStack.io, Inc.

package param

// CleanQueueDetailParam CleanQueue detail param
type CleanQueueDetailParam struct {
	SignatureName string `json:"signatureName" validate:"required"`
	TaskIndex int `json:"taskIndex,omitempty"`
	IsCleanUp bool `json:"isCleanUp,omitempty"`
	IsRunningTask bool `json:"isRunningTask,omitempty"`
	ManagementiUuid string `json:"managementiUuid,omitempty"`
}

// CleanQueueParam CleanQueue request param
type CleanQueueParam struct {
	BaseParam
	Params CleanQueueDetailParam `json:"params"`
}
