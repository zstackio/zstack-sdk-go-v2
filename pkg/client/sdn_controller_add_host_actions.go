// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SdnControllerAddHost 操作SdnControllerAddHost
func (cli *ZSClient) SdnControllerAddHost(params param.SdnControllerAddHostParam) (*view.SdnControllerAddHostEventView, error) {
	resp := view.SdnControllerAddHostEventView{}
	if err := cli.Post("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

