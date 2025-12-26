// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// EnableCdpTask operates on EnableCdpTask
func (cli *ZSClient) EnableCdpTask(params param.EnableCdpTaskParam) (*view.EnableCdpTaskEventView, error) {
	resp := view.EnableCdpTaskEventView{}
	if err := cli.Post("v1/cdp-task/enable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
