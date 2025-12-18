// Copyright (c) ZStack.io, Inc.

package param

// GetElaborationCategoriesDetailParam GetElaborationCategories detail param
type GetElaborationCategoriesDetailParam struct {
}

// GetElaborationCategoriesParam GetElaborationCategories request param
type GetElaborationCategoriesParam struct {
	BaseParam
	Params GetElaborationCategoriesDetailParam `json:"params"`
}
