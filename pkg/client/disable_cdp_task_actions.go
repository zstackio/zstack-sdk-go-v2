// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DisableCdpTask operates on DisableCdpTask
func (cli *ZSClient) DisableCdpTask(params param.DisableCdpTaskParam) (*view.DisableCdpTaskEventView, error) {
	resp := view.DisableCdpTaskEventView{}
	if err := cli.Post("v1/cdp-task/disable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
