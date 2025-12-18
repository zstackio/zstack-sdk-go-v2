// Copyright (c) ZStack.io, Inc.

package param

// GetExternalServicesDetailParam GetExternalServices detail param
type GetExternalServicesDetailParam struct {
}

// GetExternalServicesParam GetExternalServices request param
type GetExternalServicesParam struct {
	BaseParam
	Params GetExternalServicesDetailParam `json:"params"`
}
