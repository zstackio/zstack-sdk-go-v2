// Copyright (c) ZStack.io, Inc.

package param

// TakeVmConsoleScreenshotDetailParam TakeVmConsoleScreenshot detail param
type TakeVmConsoleScreenshotDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// TakeVmConsoleScreenshotParam TakeVmConsoleScreenshot request param
type TakeVmConsoleScreenshotParam struct {
	BaseParam
	Params TakeVmConsoleScreenshotDetailParam `json:"params"`
}
