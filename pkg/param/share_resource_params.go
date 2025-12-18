// Copyright (c) ZStack.io, Inc.

package param

// ShareResourceDetailParam ShareResource detail param
type ShareResourceDetailParam struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	AccountUuids []string `json:"accountUuids,omitempty"`
	ToPublic bool `json:"toPublic,omitempty"`
}

// ShareResourceParam ShareResource request param
type ShareResourceParam struct {
	BaseParam
	Params ShareResourceDetailParam `json:"params"`
}
