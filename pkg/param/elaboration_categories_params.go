// Copyright (c) ZStack.io, Inc.

package param

// GetElaborationCategoriesDetailParam GetElaborationCategories详细参数
type GetElaborationCategoriesDetailParam struct {
}

// GetElaborationCategoriesParam GetElaborationCategories请求参数
type GetElaborationCategoriesParam struct {
	BaseParam
	Params GetElaborationCategoriesDetailParam `json:"params"` // 详细参数
}

