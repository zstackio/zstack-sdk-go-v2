// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunKeySecretDetailParam AddAliyunKeySecret detail param
type AddAliyunKeySecretDetailParam struct {
	Name string `json:"name" validate:"required"`
	Key string `json:"key" validate:"required"`
	Secret string `json:"secret" validate:"required"`
	AccountUuid string `json:"accountUuid,omitempty"`
	Description string `json:"description,omitempty"`
	Sync bool `json:"sync,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunKeySecretParam AddAliyunKeySecret request param
type AddAliyunKeySecretParam struct {
	BaseParam
	Params AddAliyunKeySecretDetailParam `json:"params"`
}
