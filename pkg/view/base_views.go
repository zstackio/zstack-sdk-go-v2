// Copyright (c) ZStack.io, Inc.

package view

import (
	"time"
)

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

// Generic wrapper types for APIs that return simple data types

// MapView wraps map return values
type MapView map[string]interface{}

// ListView wraps list/array return values
type ListView []interface{}

// StringView wraps string return values
type StringView string

// BooleanView wraps boolean return values
type BooleanView bool

// IntView wraps integer return values
type IntView int

// LongView wraps long integer return values
type LongView int64

// SuccessView represents successful operation with no data return
type SuccessView struct {
	Success bool `json:"success"`
}
