// Copyright (c) ZStack.io, Inc.

package param

// AddMonToCephPrimaryStorageDetailParam AddMonToCephPrimaryStorage detail param
type AddMonToCephPrimaryStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	MonUrls []string `json:"monUrls" validate:"required"`
}

// AddMonToCephPrimaryStorageParam AddMonToCephPrimaryStorage request param
type AddMonToCephPrimaryStorageParam struct {
	BaseParam
	Params AddMonToCephPrimaryStorageDetailParam `json:"params"`
}
