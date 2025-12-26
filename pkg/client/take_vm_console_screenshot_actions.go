// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// TakeVmConsoleScreenshot operates on TakeVmConsoleScreenshot
func (cli *ZSClient) TakeVmConsoleScreenshot(uuid string, params param.TakeVmConsoleScreenshotParam) (*view.TakeVmConsoleScreenshotEventView, error) {
	resp := view.TakeVmConsoleScreenshotEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
