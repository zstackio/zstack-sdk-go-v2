// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateOssBucketParamDetail UpdateOssBucket detail param
type UpdateOssBucketParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
	OssDomain string `json:"ossDomain,omitempty"`
	OssKey string `json:"ossKey,omitempty"`
	OssSecret string `json:"ossSecret,omitempty"`
}

// UpdateOssBucketParam UpdateOssBucket request param
type UpdateOssBucketParam struct {
	BaseParam
	UpdateOssBucket UpdateOssBucketParamDetail `json:"updateOssBucket"`
}
