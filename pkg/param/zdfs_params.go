// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// ReconnectZdfsParamDetail ReconnectZdfs detail param
type ReconnectZdfsParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectZdfsParam ReconnectZdfs request param
type ReconnectZdfsParam struct {
	BaseParam
	ReconnectZdfs ReconnectZdfsParamDetail `json:"reconnectZdfs"`
}
