// Copyright (c) ZStack.io, Inc.

package param

// ResumeLongJobDetailParam ResumeLongJob detail param
type ResumeLongJobDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ResumeLongJobParam ResumeLongJob request param
type ResumeLongJobParam struct {
	BaseParam
	Params ResumeLongJobDetailParam `json:"params"`
}
