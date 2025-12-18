// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BaseInfoView base info view
type BaseInfoView struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// BaseTimeView time info view
type BaseTimeView struct {
	CreateDate time.Time `json:"createDate"`
	LastOpDate time.Time `json:"lastOpDate"`
}

// BaseResourceView resource base view
type BaseResourceView struct {
	BaseInfoView
	BaseTimeView
}
