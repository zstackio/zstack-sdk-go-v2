// Copyright (c) ZStack.io, Inc.

package param

// ExpungeImageGroupDetailParam ExpungeImageGroup detail param
type ExpungeImageGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ExpungeImageGroupParam ExpungeImageGroup request param
type ExpungeImageGroupParam struct {
	BaseParam
	Params ExpungeImageGroupDetailParam `json:"params"`
}
