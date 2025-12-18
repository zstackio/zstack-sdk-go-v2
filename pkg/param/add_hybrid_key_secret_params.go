// Copyright (c) ZStack.io, Inc.

package param

// AddHybridKeySecretDetailParam AddHybridKeySecret detail param
type AddHybridKeySecretDetailParam struct {
	Name string `json:"name" validate:"required"`
	Key string `json:"key" validate:"required"`
	Secret string `json:"secret" validate:"required"`
	AccountUuid string `json:"accountUuid,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	Sync bool `json:"sync,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddHybridKeySecretParam AddHybridKeySecret request param
type AddHybridKeySecretParam struct {
	BaseParam
	Params AddHybridKeySecretDetailParam `json:"params"`
}
