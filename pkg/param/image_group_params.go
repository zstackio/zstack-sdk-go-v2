// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// ExpungeImageGroupParamDetail ExpungeImageGroup detail param
type ExpungeImageGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ExpungeImageGroupParam ExpungeImageGroup request param
type ExpungeImageGroupParam struct {
	BaseParam
	Params ExpungeImageGroupParamDetail `json:"params"`
}
