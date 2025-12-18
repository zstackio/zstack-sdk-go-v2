// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SdnControllerChangeHost 操作SdnControllerChangeHost
func (cli *ZSClient) SdnControllerChangeHost(uuid string, params param.SdnControllerChangeHostParam) (*view.SdnControllerChangeHostEventView, error) {
	resp := view.SdnControllerChangeHostEventView{}
	if err := cli.Put("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

