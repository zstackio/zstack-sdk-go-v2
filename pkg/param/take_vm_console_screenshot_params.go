// Copyright (c) ZStack.io, Inc.

package param

// TakeVmConsoleScreenshotDetailParam TakeVmConsoleScreenshot详细参数
type TakeVmConsoleScreenshotDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// TakeVmConsoleScreenshotParam TakeVmConsoleScreenshot请求参数
type TakeVmConsoleScreenshotParam struct {
	BaseParam
	Params TakeVmConsoleScreenshotDetailParam `json:"params"` // 详细参数
}

