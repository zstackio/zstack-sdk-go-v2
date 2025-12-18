// Copyright (c) ZStack.io, Inc.

package param

// PullHuaweiIMasterControllerDetailParam PullHuaweiIMasterController detail param
type PullHuaweiIMasterControllerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	PullSwitch bool `json:"pullSwitch,omitempty"`
}

// PullHuaweiIMasterControllerParam PullHuaweiIMasterController request param
type PullHuaweiIMasterControllerParam struct {
	BaseParam
	Params PullHuaweiIMasterControllerDetailParam `json:"params"`
}
