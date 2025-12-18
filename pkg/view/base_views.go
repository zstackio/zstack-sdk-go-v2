// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BaseInfoView 基础信息视图
type BaseInfoView struct {
	UUID        string `json:"uuid"`        // 资源UUID，唯一标识
	Name        string `json:"name"`        // 资源名称
	Description string `json:"description"` // 资源描述
}

// BaseTimeView 时间信息视图
type BaseTimeView struct {
	CreateDate time.Time `json:"createDate"` // 创建时间
	LastOpDate time.Time `json:"lastOpDate"` // 最后操作时间
}

// BaseResourceView 资源基础视图，包含通用字段
type BaseResourceView struct {
	BaseInfoView
	BaseTimeView
}
