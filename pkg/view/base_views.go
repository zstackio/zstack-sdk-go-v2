// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BaseInfoView 基础信息视图（仅包含通用标识字段）
type BaseInfoView struct {
	UUID string `json:"uuid"`           // 资源唯一标识
	Name string `json:"name,omitempty"` // 资源名称
}

type BaseTimeView struct {
	CreateDate time.Time `json:"createDate,omitempty"` // 创建时间
	LastOpDate time.Time `json:"lastOpDate,omitempty"` // 最后操作时间
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
