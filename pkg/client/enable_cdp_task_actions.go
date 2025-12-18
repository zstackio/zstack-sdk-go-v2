// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// EnableCdpTask operates on EnableCdpTask
func (cli *ZSClient) EnableCdpTask(params param.EnableCdpTaskParam) (*view.EnableCdpTaskEventView, error) {
	resp := view.EnableCdpTaskEventView{}
	if err := cli.Post("v1/cdp-task/enable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
